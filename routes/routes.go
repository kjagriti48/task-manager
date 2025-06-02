package routes

import (
	"net/http"
	"task-manager/handlers"
	"task-manager/utils"
)

func RegisterRoutes() {
	http.HandleFunc("/api/signup", utils.EnableCORS(handlers.SignupHandler))
	http.HandleFunc("/api/login", utils.EnableCORS(handlers.LoginHandler))
	http.HandleFunc("/api/tasks", utils.EnableCORS(utils.JWTMiddleware(handlers.AddTaskHandler)))   // POST
	http.HandleFunc("/api/tasks/", utils.EnableCORS(utils.JWTMiddleware(handlers.ListTaskHandler))) // GET
	http.HandleFunc("/api/tasks/delete", utils.EnableCORS(utils.JWTMiddleware(handlers.DeleteTaskHandler)))
	http.HandleFunc("/api/tasks/toggle", utils.EnableCORS(utils.JWTMiddleware(handlers.ToggleTaskHandler)))

}
