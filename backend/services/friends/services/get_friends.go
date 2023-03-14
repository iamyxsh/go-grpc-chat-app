package services

import (
	"fmt"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"golang.org/x/net/context"
)

const GET_QUERY = `
SELECT  friendships.id,
		friendships.created_at,
		u1.name AS from_name,
		u1.id AS from_id,
        u2.id AS to_id,
        u2.name AS to_name
			FROM friendships
				LEFT JOIN friend_users AS u1 ON friendships.from_id = u1.id
                LEFT JOIN friend_users AS u2 ON friendships.to_id = u2.id
            WHERE friendships.status = 'accepted' AND
                  friendships.to_id = ? OR friendships.from_id = ?;
                  
`

func (s *Server) GetAllFriends(ctx context.Context, req *pb.GetAllFriendsRequest) (*pb.GetAllFriendsResponse, error) {
	friendships := []struct {
		ID        uint
		CreatedAt time.Time
		FromName  string
		FromId    uint
		ToName    string
		ToId      uint
	}{}
	err := s.DB.Raw(GET_QUERY, req.User, req.User).Scan(&friendships).Error
	if err != nil {
		fmt.Println("get requests", err.Error())
		return &pb.GetAllFriendsResponse{}, err
	}
	result := []*pb.User{}

	for _, friendship := range friendships {
		if friendship.FromId == uint(req.User) {
			result = append(result, &pb.User{
				Id:   uint32(friendship.ToId),
				Name: friendship.ToName,
			})
		} else {
			result = append(result, &pb.User{
				Id:   uint32(friendship.FromId),
				Name: friendship.FromName,
			})
		}
	}

	return &pb.GetAllFriendsResponse{Friends: result}, nil
}
