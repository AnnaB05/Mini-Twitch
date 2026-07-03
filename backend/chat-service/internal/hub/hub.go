package hub

type Hub struct {
	//all connected clients
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

func (h *Hub) Run() {
	for {
		select {
		//new client connected
		case client := <-h.register:
			h.clients[client] = true

			//client disconnected
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		//broadcast message to all clients
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					//remove client if stuck
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
