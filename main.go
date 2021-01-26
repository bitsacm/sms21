package main

import (
	"log"
	"net/http"

	"github.com/dush-t/sms21/db/models"
)

// Env will store all the stuff that the app needs
// throughout (like database access entities)
type Env struct {
	models models.Models
}

func main() {
	env := Init()
	router := Router(env)

	http.Handle("/", router)

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
