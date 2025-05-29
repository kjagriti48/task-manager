package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-manager/models"
	"task-manager/utils"
	"time"

	"github.com/google/uuid"
)

var taskStore = map[string][]models.Task{} //user -> tasks

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🚨 AddTaskHandler hit")

	usernameVal := r.Context().Value(utils.UsernameKey)
	username, ok := usernameVal.(string)

	if !ok {
		fmt.Println("❌ Username missing from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("✅ Username:", username)
	var t models.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println("❌ JSON decode error:", err)
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t.ID = uuid.New().String()
	t.User = username
	t.CreatedAt = time.Now().Format(time.RFC3339)

	taskStore[username] = append(taskStore[username], t)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)

}

func ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	usernameVal := r.Context().Value("username")
	username, ok := usernameVal.(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tasks := taskStore[username]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
