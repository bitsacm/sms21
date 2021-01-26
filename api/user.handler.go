package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dush-t/sms21/db/models"
)

// SignUpHandler is the handler function for requests at /sign_up
func SignUpHandler(data models.Models) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		addErr := data.Users.Add(user)
		if addErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error creating user:", err)
			return
		}

		tokenString, err := user.GenerateJWT()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error generating token for user:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		payload := struct {
			Token string `json:"token"`
		}{Token: tokenString}
		json.NewEncoder(w).Encode(payload)
	})
}
