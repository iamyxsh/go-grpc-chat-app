package service

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/cassandra"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	database "github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/db"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/pb"
)

type Channel struct {
	User   uint32
	Stream pb.ChatService_BroadcastMessageServer
}

func (s *Server) BroadcastMessage(stream pb.ChatService_BroadcastMessageServer) error {
	for {
		ts := (time.Now().Unix())

		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				log.Printf("Error in receiving message from client :: %v", err)
				return nil
			}
		}

		channel := Channel{
			User:   msg.From,
			Stream: stream,
		}

		if msg.Message == "" {
			if ChanMap[msg.From].Stream == nil {
				ChanMap[msg.From] = channel

				iter := s.DB.Session.Query("SELECT * FROM public.messages WHERE reciever=? AND delivered=FALSE ALLOW FILTERING;", msg.From).Iter()
				m := map[string]interface{}{}
				for iter.MapScan(m) {
					fmt.Println("rtan")
					stream.Send(&pb.Message{
						From:      uint32(m["sender"].(int)),
						To:        uint32(m["reciever"].(int)),
						Timestamp: int64(m["timestamp"].(int)),
						Message:   m["message"].(string),
					})

					err = s.DB.Session.Query("UPDATE public.messages SET delivered=? WHERE id =?", true, m["id"]).Exec()
					if err != nil {
						fmt.Println(err.Error())
					}
				}

				stream.Send(&pb.Message{Message: "Connected"})
			} else {
				delete(ChanMap, msg.From)
				return nil
			}
		} else {
			m := models.Message{
				Id:        uuid.New().String(),
				Sender:    uint(msg.From),
				Reciever:  uint(msg.To),
				Message:   msg.Message,
				Timestamp: uint(ts),
				Delivered: false,
			}

			to := ChanMap[msg.To]

			if to.User != 0 {
				m.Delivered = true
				err = to.Stream.Send(msg)
				if err != nil {
					log.Println(err)
				}
				stream.Send(&pb.Message{Message: "Delivered"})
			} else {
				stream.Send(&pb.Message{Message: "Undelivered"})
			}

			err = cassandra.InsertRow(s.DB, database.MessageTable, m)

			if err != nil {
				fmt.Println(err.Error())
			}

			if err != nil {
				log.Println(err)
			}
		}
	}
}
