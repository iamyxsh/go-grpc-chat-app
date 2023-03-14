package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/pg"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	"github.com/iamyxsh/go-grpc-chat-app/backend/services/friends/constants"
)

func ConsumeUserCreated(msg []byte) {
	user := models.FriendUser{}

	err := json.Unmarshal(msg, &user)
	if err != nil {
		fmt.Printf("error from unmarshaling user created msg %v", err.Error())
	}

	db := pg.ReturnDB(constants.DSN)
	db.Create(&user)
}
