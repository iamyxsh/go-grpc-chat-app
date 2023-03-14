package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consume(topic KAFKA_TOPICS, broker string, callback func([]byte)) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    string(topic),
		MinBytes: 0,
		MaxBytes: 10e6,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		callback(m.Value)
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
