package main

import (
	"log"
	"net"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/pg"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/constants"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/consumer"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/pb"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/service"
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

	s := grpc.NewServer()

	db := pg.ReturnDB(constants.DSN)
	db.AutoMigrate(&models.ProfileUser{})

	pb.RegisterProfileServiceServer(s, &service.Server{
		DB: db,
	})

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
