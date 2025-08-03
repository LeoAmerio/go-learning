package handlers

import (
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(rw, "Users not found")
	} else {
		models.SendData(rw, users)
	}
}

func GetUserById(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("RW:", rw)
	if user, err := getUserByReq(r); err != nil {
		models.SendNotFound(rw, "User not found")
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener el registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw, "Invalid request payload")
	} else {
		user.Save()
		models.SendData(rw, user)
	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener el registro
	var userId int64
	if user, err := getUserByReq(r); err != nil {
		models.SendNotFound(rw, "User not found")
	} else {
		userId = user.ID
	}

	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw, "Invalid request payload")
		return
	} else {
		user.ID = userId
		user.Save()
		models.SendData(rw, user)
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByReq(r); err != nil {
		models.SendNotFound(rw, "User not found")
	} else {
		user.Delete()
		models.SendData(rw, "User deleted successfully")
	}
}

func getUserByReq(r *http.Request) (*models.User, error) {
	vars := mux.Vars(r)
	userIdStr, exists := vars["id"]
	if !exists {
		return nil, fmt.Errorf("user ID not provided in URL")
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	return models.GetUserByID(userId)
}
