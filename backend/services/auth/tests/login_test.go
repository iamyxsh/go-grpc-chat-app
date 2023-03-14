package tests

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/iamyxsh/go-grpc-chat-app/backend/services/auth/pb"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {
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

	client.Signup(context.Background(), &pb.SignupRequest{
		Email:    EMAIL,
		Name:     NAME,
		Password: PASSWORD,
	})

	res, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    EMAIL,
		Password: PASSWORD,
	})
	if err != nil {
		clearDB(db)
		t.Fatalf("signup user %v", err)
	}

	assert.True(strings.HasPrefix(res.Msg, TOKEN_PREFIX), "token is malformed")
}
