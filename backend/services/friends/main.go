package main

import (
	"log"
	"net"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/pg"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/utils"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/constants"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/consumer"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/database"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/services"
	"google.golang.org/grpc"
)

func main() {
	go kafka.Consume(kafka.USER_CREATED, kafka.BROKER, consumer.ConsumeUserCreated)
	lis, err := net.Listen("tcp", constants.Addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", constants.Addr)

	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.ChainUnaryInterceptor(utils.CheckHeaderInterceptor()))

	s := grpc.NewServer(opts...)
	defer s.Stop()

	db := pg.ReturnDB(constants.DSN)
	database.ExecStatements(db)

	pb.RegisterFriendsServiceServer(s, &services.Server{
		DB: db,
	})
 
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
} 
    