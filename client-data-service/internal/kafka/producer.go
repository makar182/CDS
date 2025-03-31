package kafka

import (
	"github.com/segmentio/kafka-go"
	"log"
)

func NewProducer() (*kafka.Writer, error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "logs",
		Balancer: &kafka.LeastBytes{},
	}

	log.Println("Kafka producer инициализирован")
	return writer, nil
}
