package constants

import (
	"fmt"
	"os"
)

const DSN = "postgres://postgres:postgres@friends-db:5432/postgres"

var Addr string = fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))
