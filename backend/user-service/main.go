package main

import (
	"log"
	"net/http"

	"user-service/cmd/api"
)

func main() {
	router := api.NewRouter()

	log.Println("User Service running on :8080")
	http.ListenAndServe(":8080", router)
}
