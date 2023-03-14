package database

import (
	"log"
	"os"
	"strconv"

	"github.com/iamyxsh/go-grpc-chat-app/backend/packages/db/cassandra"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	messageTableName = "public.messages"
	userTableName    = "public.users"
)

var (
	MessageTable table.Table
	UserTable    table.Table
)

func init() {
	p, _ := strconv.Atoi(os.Getenv("CASSANDRA_PORT"))
	session := cassandra.ReturnDB(os.Getenv("CASSANDRA_HOST"), p)
	err := session.ExecStmt(`CREATE TABLE IF NOT EXISTS public.messages (
		id varchar PRIMARY KEY,
		sender int,
		reciever int,
		message text,
		timestamp int, 
		delivered boolean
	)`)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS public.users (
		id int PRIMARY KEY,
		name text
	)`)
	if err != nil {
		log.Fatalf(err.Error())
	}

	msgMetadata := cassandra.CreateMetadata(messageTableName, []string{"id", "sender", "reciever", "timestamp", "message", "delivered"}, []string{"id"})
	msgTable := cassandra.CreateTable(msgMetadata)

	userMetadata := cassandra.CreateMetadata(userTableName, []string{"id", "name"}, []string{"id"})
	userTable := cassandra.CreateTable(userMetadata)

	MessageTable = *msgTable
	UserTable = *userTable
}
