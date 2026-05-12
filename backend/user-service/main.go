package main

import (
	"log"      //log package
	"net/http" //http packag

	"user-service/cmd/api" //api package from user-service module
)

func main() {
	router := api.NewRouter()

	log.Println("User Service running on :8080")
	http.ListenAndServe(":8080", router) //starts http server on port 8080
}
