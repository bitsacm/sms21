package api

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/dush-t/sms21/util"
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

// SignInHandler is the handler function for requests at /sign_in
func SignInHandler(data models.Models) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if(err != nil) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, readErr := data.Users.GetUserByUsername(body.Username)

		if readErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error reading user:", err)
			return
		}

		passwordMatch := util.MatchesWithHash(body.Password, user.Password)

		if (user == models.User{} || passwordMatch != true) {
			w.WriteHeader(http.StatusUnauthorized)
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
		payload := map[string]interface{} {
			"token": tokenString,
			"user": struct {
				ID string `json:"id"`
				Username string `json:"username"`
				Name string `json:"name"`
				RegToken string `json:"regToken"`
			} {user.ID, user.Username, user.Name, user.RegToken},
		}
		json.NewEncoder(w).Encode(payload)
	})
}
