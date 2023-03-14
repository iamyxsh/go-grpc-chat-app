package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/utils"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	user := models.AuthUser{
		Email: req.Email,
	}

	res := s.DB.Where("email = ?", req.Email).First(&user)

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.Unknown, res.Error.Error())
	}

	err := utils.Hash{}.CompareHash(req.Password, user.Password)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Unknown, res.Error.Error())
	}

	token, err := utils.JWT{}.GenerateToken(strconv.FormatUint(uint64(user.ID), 32), "secret", time.Now().Add(time.Hour*time.Duration(24)))
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.Response{Msg: token}, nil
}
