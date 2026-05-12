package main

import (
	"log"
	"net/http"

	"auth-service/cmd/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Auth Service running on :8081")
	http.ListenAndServe(":8081", router)
}
