package api

import (
	"net/http" //http package

	"github.com/go-chi/chi/v5" //chi router package
)

func NewRouter() http.Handler {
	r := chi.NewRouter() //creates chi router instance

	//health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user-service OK"))
	})

	return r
}
