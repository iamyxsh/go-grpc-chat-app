package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/cassandra"
	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/models"
	database "github.com/iamyxsh/go-grpc-chat-app/backend/services/chat/db"
)

func ConsumeUserCreated(msg []byte) {
	user := models.FriendUser{}

	err := json.Unmarshal(msg, &user)
	if err != nil {
		fmt.Printf("error from unmarshaling user created msg %v", err.Error())
	}

	p, _ := strconv.Atoi(os.Getenv("CASSANDRA_PORT"))
	session := cassandra.ReturnDB(os.Getenv("CASSANDRA_HOST"), p)

	err = cassandra.InsertRow(session, database.UserTable, user)
	if err != nil {
		log.Fatal(err.Error())
	}
}
