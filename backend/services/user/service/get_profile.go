package service

import (
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	user := models.ProfileUser{}
	err := s.DB.Find(&user, req.Id).Error
	if err != nil {
		return &pb.GetProfileResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetProfileResponse{Email: user.Email, Name: user.Name, Id: uint32(user.ID)}, nil
}
