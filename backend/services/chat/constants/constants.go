package constants

import (
	"fmt"
	"os"
)

var Addr string = fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))
