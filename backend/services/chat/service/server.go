package service

import (
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/pb"
	"github.com/scylladb/gocqlx/v2"
)

type Server struct {
	DB gocqlx.Session
	// Logger *logrus.Logger
	pb.UnimplementedChatServiceServer
}

var ChanMap = make(map[uint32]Channel)
