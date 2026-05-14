package main

import (
	"log"
	"net/http"

	"auth-service/cmd/api"
	"auth-service/internal/database"
)

func main() {
	db, err := database.Connect(
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"authdb",
	)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Connected to PostgreSQL")
	router := api.NewRouter()

	log.Println("Auth Service running on :8081")
	http.ListenAndServe(":8081", router)
	_ = db
}
