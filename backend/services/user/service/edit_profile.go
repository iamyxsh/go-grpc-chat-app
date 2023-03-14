package service

import (
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) EditProfile(ctx context.Context, req *pb.EditProfileRequest) (*pb.Response, error) {
	user := models.ProfileUser{}
	err := s.DB.Find(&user, req.Id).Error
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	err = s.DB.Model(&models.ProfileUser{}).Where("id = ?", req.Id).Update("name", req.Name).Error
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.Response{Msg: "profile updated"}, nil
}
