package kafka

import (
	"context"
	"fmt"
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

func Write() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "messages", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(kafka.Message{Value: []byte("test")})
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func Read() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "messages", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
