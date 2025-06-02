package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var userStore = map[string]string{}  // username -> hashed password
var jwtKey = []byte("my-secret-key") // ✅ Must be the exact same string

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}
	if u.Email == "" || u.Password == "" {
		writeError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	if _, exists := userStore[u.Email]; exists {
		writeError(w, http.StatusBadRequest, "User already exits")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	userStore[u.Email] = string(hashed)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	hashedPassword, ok := userStore[u.Email]
	if !ok {
		writeError(w, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		writeError(w, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not sign token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenStr,
	})
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
