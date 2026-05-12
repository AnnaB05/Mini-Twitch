package main

import (
	"log"
	"net/http"

	"chat-service/cmd/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Chat Service running on :8082")
	http.ListenAndServe(":8082", router)
}
