package main

import (
	"client-data-service/internal/database"
	"client-data-service/internal/handlers"
	"client-data-service/internal/kafka"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}
	defer db.Close()

	// Инициализация Kafka producer
	producer, err := kafka.NewProducer()
	if err != nil {
		log.Fatalf("Ошибка инициализации Kafka producer: %v", err)
	}
	defer producer.Close()

	// HTTP маршруты
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет, как дела?")
	})

	http.HandleFunc("/submit", handlers.SubmitHandler(db, producer))

	log.Println("Сервис управления клиентскими данными запущен на :8080")

	// Запуск HTTP-сервера
	log.Fatal(http.ListenAndServe(":8080", nil))
}
