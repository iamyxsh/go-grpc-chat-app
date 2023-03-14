package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type ENV string

const (
	PROD ENV = "production"
	DEV  ENV = "development"
)

func ReturnLogger(env ENV) {
	if env == PROD {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.DebugLevel)

		file, err := os.OpenFile("logger/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	} else {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.ErrorLevel)
		log.SetOutput(os.Stdout)
	}
}
