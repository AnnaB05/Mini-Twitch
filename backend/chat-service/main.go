package main

import (
	"log"
	"net/http"

	"chat-service/cmd/api"
	"chat-service/internal/hub"
	WS "chat-service/internal/websocket"
)

func main() {

	//create the hub
	hub := hub.NewHub()

	//run the hub
	go hub.Run()

	//create router
	router := api.NewRouter()

	//add websocket route to router
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WS.ServeWs(hub, w, r)
	})
	
	//start the server
	log.Println("Chat Service running on :8082")
	http.ListenAndServe(":8082", router)
}
