package main

import (
	"apirest/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/user", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUserById).Methods("GET")
	mux.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("🚀 Server starting on http://localhost:8080")
	fmt.Println("📋 Available endpoints:")
	fmt.Println("  GET    /api/user       - List all users")
	fmt.Println("  GET    /api/user/{id}  - Get user by ID")
	fmt.Println("  POST   /api/user       - Create new user")
	fmt.Println("  PUT    /api/user/{id}  - Update user")
	fmt.Println("  DELETE /api/user/{id}  - Delete user")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
