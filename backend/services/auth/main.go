package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/pg"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/service"
	"google.golang.org/grpc"
)

const DSN = "postgres://postgres:postgres@auth-db:5432/postgres"

var addr string = fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()

	db := pg.ReturnDB(DSN)
	db.AutoMigrate(&models.AuthUser{})

	pb.RegisterAuthServiceServer(s, &service.Server{
		DB: db,
	})

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
 