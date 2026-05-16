package handlers

import (
	"encoding/json"
	"net/http"

	"auth-service/cmd/api/middleware"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUser(r)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": user.UserID,
		"email":   user.Email,
	})
}
