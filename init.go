package main

import (
	"log"

	"github.com/dush-t/sms21/db"
	"github.com/dush-t/sms21/db/models"
)

// Init does all the setup work to start the app
// (like creating a database connection)
func Init() Env {
	const (
		dbURI      string = "neo4j://192.168.2.2:7687"
		dbUsername string = "neo4j"
		dbPassword string = "test"
	)

	conn, err := db.NewConn(dbURI, dbUsername, dbPassword)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	models := models.Init(conn)
	return Env{
		models: *models,
	}
}
