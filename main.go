package main

import (
	"fmt"
	"log"
	"net/http"
	"task-manager/routes"
)

func main() {
	routes.RegisterRoutes()

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
