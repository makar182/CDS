package kafka

import (
	"context"
	"database/sql"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	db     *sql.DB
}

func NewConsumer(db *sql.DB) (*Consumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "logs",
		GroupID: "log-consumer-group",
	})

	return &Consumer{reader: reader, db: db}, nil
}

func (c *Consumer) ConsumeMessages() {
	log.Println("Начало чтения логов из Kafka...")

	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v", err)
			continue
		}

		// Сохранение в БД
		_, err = c.db.Exec("INSERT INTO logs (message) VALUES ($1)", string(m.Value))
		if err != nil {
			log.Printf("Ошибка записи лога в БД: %v", err)
		}
	}
}
