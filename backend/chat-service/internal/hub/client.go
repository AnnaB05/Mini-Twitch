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

func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		//send message to the hub for broadcasting
		c.hub.broadcast <- message
	}
}
