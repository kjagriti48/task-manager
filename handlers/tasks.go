package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("❌ Could not read body:", err)
		http.Error(w, "Could not read body", http.StatusInternalServerError)
		return
	}

	fmt.Println("📦 Raw body:", string(bodyBytes))
	r.Body = io.NopCloser(strings.NewReader(string(bodyBytes))) // reset body reader

	err = json.NewDecoder(r.Body).Decode(&t) // ✅ just use = instead of :=
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

	fmt.Println("✅ Task stored:", t)
	fmt.Println("🧠 Current user task list:", taskStore[username])
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)

	fmt.Println("✅ Storing task in memory:", t)
	fmt.Println("🧠 Current tasks for user:", taskStore[username])

}

func ListTaskHandler(w http.ResponseWriter, r *http.Request) {
	usernameVal := r.Context().Value(utils.UsernameKey)
	username, ok := usernameVal.(string)
	if !ok || username == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Println("✅ Username from context:", username)

	// 💡 Return the real user-specific task list
	tasks := taskStore[username]
	for user, tasks := range taskStore {
		fmt.Printf("📦 All tasks for %s: %v\n", user, tasks)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(utils.UsernameKey).(string)
	id := r.URL.Query().Get("id")

	tasks := taskStore[username]
	filtered := []models.Task{}
	for _, t := range tasks {
		if t.ID != id {
			filtered = append(filtered, t)
		}
	}
	taskStore[username] = filtered
	w.WriteHeader(http.StatusNoContent)
}

func ToggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(utils.UsernameKey).(string)
	id := r.URL.Query().Get("id")

	for i, t := range taskStore[username] {
		if t.ID == id {
			taskStore[username][i].Completed = !t.Completed
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}
