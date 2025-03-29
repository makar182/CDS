package cmd

import (
	"log"
	"net/http"

	"logging-service/internal/database"
	"logging-service/internal/handlers"
	"logging-service/internal/kafka"
)

func Main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}
	defer db.Close()

	// Инициализация Kafka consumer
	consumer, err := kafka.NewConsumer(db)
	if err != nil {
		log.Fatalf("Ошибка инициализации Kafka consumer: %v", err)
	}
	go consumer.ConsumeMessages()

	// HTTP маршруты
	http.HandleFunc("/logs", handlers.GetLogsHandler(db))

	log.Println("Сервис логирования запущен на :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
