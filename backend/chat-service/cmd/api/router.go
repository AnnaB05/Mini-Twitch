package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"chat-service/internal/hub"
	ws "chat-service/internal/websocket"
)

func NewRouter(h *hub.Hub) http.Handler {
	r := chi.NewRouter()

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("chat-service OK"))
	})

	//websocket route
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(h, w, r)
	})

	return r
}
