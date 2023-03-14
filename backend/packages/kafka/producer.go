package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Produce(topic KAFKA_TOPICS, broker, key string, msg []byte) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    string(topic),
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: msg,
		},
	)
	if err != nil {
		fmt.Println("producer error", err.Error())
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return nil
}
