package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dush-t/sms21/db/models"
)

func AddTransactionHandler(data models.Models) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction

		err := json.NewDecoder(r.Body).Decode(&transaction)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, addErr := data.Transactions.Add(transaction)
		if addErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error creating stock:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		payload := struct {
			ID string `json:"id"`
		}{id}

		json.NewEncoder(w).Encode(payload)
	})
}
