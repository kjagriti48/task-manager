package routes

import (
	"net/http"
	"task-manager/handlers"
	"task-manager/utils"
)

func RegisterRoutes() {
	http.HandleFunc("/signup", utils.EnableCORS(handlers.SignupHandler))
	http.HandleFunc("/login", utils.EnableCORS(handlers.LoginHandler))
	http.HandleFunc("/tasks", utils.EnableCORS(utils.JWTMiddleware(handlers.AddTaskHandler)))   // POST
	http.HandleFunc("/tasks/", utils.EnableCORS(utils.JWTMiddleware(handlers.ListTaskHandler))) // GET

}
