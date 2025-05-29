package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UsernameKey contextKey = "username"

var jwtKey = []byte("my-secret-key") // ✅ Must be the exact same string

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println("🔍 Raw Authorization:", authHeader)
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			fmt.Println("❌ JWT validation error:", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		fmt.Println("✅ Token claims:", claims)

		username, ok := (*claims)["username"].(string)
		if !ok || username == "" {
			http.Error(w, "Username missing in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
