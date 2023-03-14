package services

import (
	"errors"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Server) RejectFriend(ctx context.Context, req *pb.RejectFriendRequest) (*pb.Response, error) {
	friendship := models.Friendship{}
	res := s.DB.First(&friendship, req.Id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.NotFound, "no requests found")
	}

	s.DB.Model(&models.Friendship{}).Where("id = ?", req.Id).Update("status", models.REJECTED)

	return &pb.Response{Msg: "friend request rejected"}, nil
}
