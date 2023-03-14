package services

import (
	"fmt"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"golang.org/x/net/context"
)

const JOIN_QUERY = `
SELECT  friendships.status,
		friendships.id,
		friendships.created_at,
		u1.name AS from_name, 
		u1.id AS from_id,
		u2.name AS to_name 
			FROM friendships
				LEFT JOIN friend_users AS u1 ON friendships.from_id = u1.id
				LEFT JOIN friend_users AS u2 ON friendships.to_id = u2.id
		WHERE friendships.to_id = ?;
`

func (s *Server) GetAllFriendRequests(ctx context.Context, req *pb.GetAllFriendRequestsRequest) (*pb.GetAllFriendRequestsResponse, error) {
	friendships := []struct {
		ID        uint
		Status    models.FriendshipStatus
		CreatedAt time.Time
		FromName  string
		FromId    uint
		ToName    string
	}{}
	err := s.DB.Raw(JOIN_QUERY, req.User).Scan(&friendships).Error
	if err != nil {
		fmt.Println("get requests", err.Error())
		return &pb.GetAllFriendRequestsResponse{Requests: []*pb.FriendRequests{}}, err
	}

	result := []*pb.FriendRequests{}

	for _, friendship := range friendships {
		result = append(result, &pb.FriendRequests{
			Id: uint32(friendship.ID),
			From: &pb.User{
				Id:   uint32(friendship.FromId),
				Name: friendship.FromName,
			},
			To: &pb.User{
				Id:   uint32(req.User),
				Name: friendship.ToName,
			},
			Time: friendship.CreatedAt.String(),
		})
	}

	return &pb.GetAllFriendRequestsResponse{Requests: result}, nil
}
