package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KAFKA_TOPICS string

const (
	USER_CREATED KAFKA_TOPICS = "USER_CREATED"
)

const BROKER = "kafka:9092"

func init() {
	controllerConn, err := kafka.Dial("tcp", BROKER)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{{Topic: string(USER_CREATED), NumPartitions: 1, ReplicationFactor: 0}}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		fmt.Println(err.Error())
	}
}
