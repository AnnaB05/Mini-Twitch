package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"auth-service/internal/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true //allow all origins
	},
}

func ServeWs(hub *hub.Hub, w http.ResponseWriter, r *http.Request) {
	
	//upgrade HTTP connection to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	//create new client
	client := &hub.Client{
		conn: conn,
		send: make(chan []byte, 256),
		hub:  hub,
	}

	//register client with hub
	client.hub.register <- client

	//start read and write pumps
	go client.WritePump()
	go client.ReadPump()
}
