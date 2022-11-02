package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type Msg struct {
	Text   string
	Target string
	From   string
}

func NewMsg(msg string, target string, from string) Msg {
	return Msg{
		Text:   msg,
		Target: target,
		From:   from,
	}
}

func Configure(url string, topic string, partition int) (c *kafka.Conn, err error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", url, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn, err
}

func Write(message Msg) {
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
	m, _ := json.Marshal(message)
	_, err = conn.WriteMessages(kafka.Message{Value: m})
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func Read(url string, topic string, callback func(Msg)) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{url},
		GroupID:  "consumer-group-id",
		Topic:    topic,
		MinBytes: 0,    // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background()) // 10KB max per message
		if err != nil {
			log.Println(err)
			break
		}
		msg := Msg{}
		err = json.Unmarshal(m.Value, &msg)
		if err != nil {
			return
		}
		log.Println(msg.Text)
		callback(msg)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
