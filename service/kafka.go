package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func Configure(url string, topic string, partition int) (c *kafka.Conn, err error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", url, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn, err
}

func Write(message string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "messages", 0)
	log.Println("write")

	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal("failed to set deadline:", err)
		return
	}
	_, err = conn.WriteMessages(kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func Read(callback func(string)) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "messages",
		MinBytes: 0,    // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background()) // 10KB max per message
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(m.Value))
		callback(string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
