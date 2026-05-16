package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"auth-service/internal/database"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// user registration handler
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest //parse JSON body

	//parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	//trim em and pw whitespace
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)

	//validate email/pw
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)

		return
	}

	//hash pw with bcrypt
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	//hash error handling
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	//get db instance
	db := database.Get()

	//add new user to db with hashed pw
	_, err = db.Conn.Exec(
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		req.Email,
		string(hashed),
	)
	//duplicate email handling
	if err != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	//201 created/JSON response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"registered"}`))
}
