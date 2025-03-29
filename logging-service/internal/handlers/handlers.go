package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetLogsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, message FROM logs ORDER BY id DESC")
		if err != nil {
			http.Error(w, "Ошибка запроса к БД", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var logs []map[string]interface{}
		for rows.Next() {
			var id int
			var message string
			if err := rows.Scan(&id, &message); err != nil {
				http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
				return
			}
			logs = append(logs, map[string]interface{}{"id": id, "message": message})
		}

		json.NewEncoder(w).Encode(logs)
	}
}
