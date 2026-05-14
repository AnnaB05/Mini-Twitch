package api

import (
	"auth-service/cmd/api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	//middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-service OK"))
	})

	//auth routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handlers.Register)
		r.Post("/login", handleLogin)
		r.Post("/refresh", handleRefresh)
		r.Post("/logout", handleLogout)

	})

	return r
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login endpoint"))
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("refresh endpoint"))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout endpoint"))
}
