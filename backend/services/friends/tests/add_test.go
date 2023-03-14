package tests

import (
	"context"
	"log"
	"testing"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/database"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/test/bufconn"
)

func TestFriendService_AddFriend(t *testing.T) {
	t.Skip()
	assert := assert.New(t)

	db := returnDB()
	clearDB(db)
	database.ExecStatements(db)
	createEntries(db)

	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()

	srv := returnServer(db)
	defer srv.Stop()

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	client, cancel := returnClient(lis)
	defer cancel()

	res, err := client.AddFriend(context.Background(), &pb.AddFriendRequest{
		To:   2,
		From: 1,
	})
	if err != nil {
		clearDB(db)
		t.Fatalf("signup user %v", err)
	}

	friendship := models.Friendship{}
	db.First(&friendship, "to_id = ? AND from_id = ?", 2, 1)

	assert.Equal(res.Msg, "friend request sent", "msg is wrong")
	assert.Equal(uint(1), friendship.ID, "id is wrong")
}
