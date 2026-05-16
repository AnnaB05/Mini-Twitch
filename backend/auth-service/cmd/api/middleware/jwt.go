package middleware

import (
	"context"
	"net/http"
	"strings"

	"auth-service/internal/jwt"
)

type contextKey string

const userContextKey = contextKey("user")

type AuthenticatedUser struct {
	UserID int
	Email  string
}

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//read Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		//must be "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		//parse + validate JWT
		claims, err := jwt.Validate(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		//attach user info to context
		user := AuthenticatedUser{
			UserID: claims.UserID,
			Email:  claims.Email,
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// get user from context
func GetUser(r *http.Request) AuthenticatedUser {
	user, _ := r.Context().Value(userContextKey).(AuthenticatedUser)
	return user
}
