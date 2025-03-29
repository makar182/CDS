package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"net/http"

	"client-data-service/internal/models"
	//"github.com/segmentio/kafka-go"
)

func SubmitHandler(db *sql.DB, producer *kafka.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
			return
		}

		// Проверка полей
		if user.FirstName == "" || user.LastName == "" || user.MiddleName == "" {
			http.Error(w, "Все поля должны быть заполнены", http.StatusBadRequest)
			return
		}

		// Сохранение в БД
		_, err := db.Exec("INSERT INTO users (first_name, last_name, middle_name) VALUES ($1, $2, $3)",
			user.FirstName, user.LastName, user.MiddleName)
		if err != nil {
			http.Error(w, "Ошибка при сохранении", http.StatusInternalServerError)
			return
		}

		// Отправка в Kafka
		message := fmt.Sprintf("Добавлен пользователь: %s %s %s", user.FirstName, user.LastName, user.MiddleName)
		err = producer.WriteMessages(r.Context(), kafka.Message{
			Value: []byte(message),
		})
		if err != nil {
			http.Error(w, "Ошибка отправки в Kafka", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
