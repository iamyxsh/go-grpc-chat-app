package services

import (
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	// Logger *logrus.Logger
	pb.UnimplementedFriendsServiceServer
}
