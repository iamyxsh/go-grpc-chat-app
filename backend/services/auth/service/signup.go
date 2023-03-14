package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/kafka"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/utils"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Server) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.Response, error) {
	user := models.AuthUser{
		Email: req.Email,
	}

	res := s.DB.Where("email = ?", req.Email).First(&user)

	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.Unknown, res.Error.Error())
	}

	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &pb.Response{}, status.Error(codes.AlreadyExists, "user already exists")
	}

	hash, err := utils.Hash{}.GenerateHash(req.Password)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Unknown, res.Error.Error())
	}
	user = models.AuthUser{
		Email:    req.Email,
		Password: hash,
	}

	user.Password = hash

	s.DB.Create(&user)

	jsonString, err := json.Marshal(struct {
		ID       uint
		Name     string
		Email    string
		Password string
	}{
		ID:       user.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
	})
	if err != nil {
		return &pb.Response{}, status.Error(codes.Unknown, res.Error.Error())
	}

	token, err := utils.JWT{}.GenerateToken(strconv.FormatUint(uint64(user.ID), 32), "secret", time.Now().Add(time.Hour*time.Duration(24)))
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	go kafka.Produce(kafka.USER_CREATED, kafka.BROKER, "user", jsonString)

	return &pb.Response{Msg: token}, nil
}
