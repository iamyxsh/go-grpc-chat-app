package service

import (
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	// Logger *logrus.Logger
	pb.UnimplementedAuthServiceServer
}
