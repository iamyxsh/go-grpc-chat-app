package tests

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	To           = "Yash"
	ToID         = 1
	From         = "Sharma"
	FromID       = 2
	DB_FILE_NAME = "test.db"
)

func clearDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.Friendship{})
	db.Migrator().DropTable(&models.FriendUser{})
}

func returnServer(db *gorm.DB) *grpc.Server {
	srv := grpc.NewServer()

	pb.RegisterFriendsServiceServer(srv, &services.Server{DB: db})

	return srv
}

func returnClient(lis *bufconn.Listener) (pb.FriendsServiceClient, context.CancelFunc) {
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	conn, _ := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))

	return pb.NewFriendsServiceClient(conn), cancel
}

func returnDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_FILE_NAME), &gorm.Config{})
	if err != nil {
		log.Fatalln("error from db", err.Error())
	}
	return db
}

func createEntries(db *gorm.DB) (*models.FriendUser, *models.FriendUser) {
	to := models.FriendUser{
		ID:   ToID,
		Name: To,
	}
	from := models.FriendUser{
		ID:   FromID,
		Name: From,
	}

	db.Create(&to)
	db.Create(&from)

	return &to, &from
}
