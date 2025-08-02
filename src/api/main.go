package main

import (
	"apirest/handlers"
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

	http.ListenAndServe(":8080", mux)

}
