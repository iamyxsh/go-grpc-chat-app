package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/cassandra"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/constants"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/consumer"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/pb"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/service"
	"google.golang.org/grpc"
)

func main() {
	go kafka.Consume(kafka.USER_CREATED, kafka.BROKER, consumer.ConsumeUserCreated)

	lis, err := net.Listen("tcp", constants.Addr)
	fmt.Println(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", constants.Addr)

	s := grpc.NewServer()
	defer s.Stop()
	p, _ := strconv.Atoi(os.Getenv("CASSANDRA_PORT"))
	db := cassandra.ReturnDB(os.Getenv("CASSANDRA_HOST"), p)

	pb.RegisterChatServiceServer(s, &service.Server{
		DB: db,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
