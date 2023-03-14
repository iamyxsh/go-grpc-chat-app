package tests

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_Signup(t *testing.T) {
	// t.Skip()
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

	res, err := client.Signup(context.Background(), &pb.SignupRequest{
		Email:    EMAIL,
		Name:     NAME,
		Password: PASSWORD,
	})
	if err != nil {
		clearDB(db)
		t.Fatalf("signup user %v", err)
	}

	user := models.AuthUser{}
	db.First(&user, "email = ?", EMAIL)

	assert.True(strings.HasPrefix(res.Msg, TOKEN_PREFIX), "token is malformed")
	assert.True(strings.HasPrefix(user.Password, HASH_PREFIX), "password is malformed")
}
