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

func (s *Server) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.Response, error) {
	to := models.FriendUser{
		ID: uint(req.To),
	}
	from := models.FriendUser{
		ID: uint(req.From),
	}

	res := s.DB.First(&to)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.NotFound, "to user does not exists")
	}
	res = s.DB.First(&from)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.NotFound, "from user does not exists")
	}

	friendship := models.Friendship{}
	res = s.DB.First(&friendship, "to_id = ? AND from_id = ?", req.To, req.From)

	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.NotFound, "request already sent")
	}

	friendship = models.Friendship{
		To:     to,
		From:   from,
		Status: models.SENT,
	}
	s.DB.Create(&friendship)

	return &pb.Response{Msg: "friend request sent"}, nil
}
