package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db.Connect()

	users := models.ListUsers()
	db.Close()

	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))

}

func GetUserById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el ID de la URL
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUserByID(int64(userId))
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validar campos requeridos
	if user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(rw, "Missing required fields: username, password, and email are required", http.StatusBadRequest)
		return
	}

	db.Connect()
	defer db.Close()

	// Crear el usuario con los datos recibidos
	createdUser := models.CreateUser(user.Username, user.Password, user.Email)
	createdUser.Save()
	fmt.Println("%s", createdUser)
	// Retornar el usuario creado
	output, err := json.Marshal(createdUser)
	if err != nil {
		http.Error(rw, "Error creating response", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write(output)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validar campos requeridos
	if user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(rw, "Missing required fields: username, password, and email are required", http.StatusBadRequest)
		return
	}

	db.Connect()
	defer db.Close()

	// Crear el usuario con los datos recibidos
	user.Save()

	// Retornar el usuario creado
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Error creating response", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write(output)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener el ID de la URL
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUserByID(int64(userId))
	if user == nil {
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}
	user.Delete()
	fmt.Printf("User with ID %d deleted successfully\n", userId)

	db.Close()
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(rw, "Error creating response", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(rw, string(output))

	rw.WriteHeader(http.StatusNoContent)
}
