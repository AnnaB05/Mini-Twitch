package api

import (
	"auth-service/cmd/api/handlers"
	"auth-service/cmd/api/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

// set up http router with chi
func NewRouter() http.Handler {
	r := chi.NewRouter()

	//middleware
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	//health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-service OK"))
	})

	//auth routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handlers.Register)
		r.Post("/login", handlers.Login)
		r.Post("/refresh", handleRefresh)
		r.Post("/logout", handleLogout)

	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.JWT)

		r.Get("/auth/me", handlers.Me) //testing protected route
	})

	return r
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("refresh endpoint"))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout endpoint"))
}
