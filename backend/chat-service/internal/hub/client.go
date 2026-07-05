package hub

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	//websocket connection
	conn *websocket.Conn

	//channel for sending messages to the client
	send chan []byte

	//reference to the hub
	hub *Hub
}
