package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/user/pb"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stretchr/testify/assert"
)

func TestProfileService_Edit(t *testing.T) {
	t.Skip()
	assert := assert.New(t)

	db := returnDB()

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

	user := models.ProfileUser{
		ID:        1,
		Email:     EMAIL,
		Name:      NAME,
		Password:  PASSWORD,
		CreatedAt: time.Now(),
	}

	db.Create(&user)

	res, err := client.EditProfile(context.Background(), &pb.EditProfileRequest{
		Name: "Rahul Sharma",
		Id:   1,
	})
	if err != nil {
		clearDB(db)
		t.Fatalf("signup user %v", err)
	}

	user1 := models.ProfileUser{}

	db.First(&user1, 1)

	assert.Equal(res.Msg, "profile updated", "msg not correct")
	assert.Equal(user1.Name, "Rahul Sharma", "name not correct")
}
