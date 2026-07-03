package hub

type Hub struct {
	//connected clients
	clients map[*Client]bool

	//channel for broadcasting messages to all clients
	broadcast chan []byte

	//channel for registering new clients
	register chan *Client

	//channel for unregistering clients
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}
