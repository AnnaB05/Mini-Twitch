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

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		message, ok := <-c.send
		if !ok {
			//hub closed the channel/client is disconnected
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}
		//write message to the websocket connection
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}

}
