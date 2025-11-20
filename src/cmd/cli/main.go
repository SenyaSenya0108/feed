package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           []string{"kafka:9092"},
		Topic:             "category",
		GroupID:           "category",
		CommitInterval:    1 * time.Second,
		ReadLagInterval:   1 * time.Second,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    10 * time.Second,
		RebalanceTimeout:  10 * time.Second,
	})
	defer reader.Close()

	log.Println("start reading message")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка чтения: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf("Получено: %s\n", string(msg.Value))
		break
	}
}
