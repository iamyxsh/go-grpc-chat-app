package tests

import (
	"context"
	"log"
	"testing"

	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/database"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/test/bufconn"
)

func TestFriendService_GetFriends(t *testing.T) {
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

	_, err := client.AddFriend(context.Background(), &pb.AddFriendRequest{
		To:   1,
		From: 2,
	})
	if err != nil {
		t.Fatalf("create friends %v", err)
	}

	_, err = client.AcceptFriend(context.Background(), &pb.AcceptFriendRequest{
		Id: 1,
	})
	if err != nil {
		t.Fatalf("create friends %v", err)
	}

	res, err := client.GetAllFriends(context.Background(), &pb.GetAllFriendsRequest{
		User: 1,
	})
	if err != nil {
		clearDB(db)
		t.Fatalf("ghet friend requests %v", err)
	}

	assert.Greater(len(res.Friends), 0, "requests should be more than 0")
	assert.Equal(uint(2), uint(res.Friends[0].Id), "id is wrong")
}
