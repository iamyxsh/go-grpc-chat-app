package tests

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	NAME         = "Yash Sharma"
	EMAIL        = "iamyxsh@icloud.com"
	PASSWORD     = "tedmosby"
	TOKEN_PREFIX = "ey"
	HASH_PREFIX  = "$2a$10$"
	DB_FILE_NAME = "test.db"
)

func clearDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.AuthUser{})
}

func returnServer(db *gorm.DB) *grpc.Server {
	srv := grpc.NewServer()

	pb.RegisterAuthServiceServer(srv, &service.Server{DB: db})

	return srv
}

func returnClient(lis *bufconn.Listener) (pb.AuthServiceClient, context.CancelFunc) {
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	conn, _ := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))

	return pb.NewAuthServiceClient(conn), cancel
}

func returnDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_FILE_NAME), &gorm.Config{})
	if err != nil {
		log.Fatalln("error from db", err.Error())
	}
	clearDB(db)
	db.AutoMigrate(&models.AuthUser{})
	return db
}
