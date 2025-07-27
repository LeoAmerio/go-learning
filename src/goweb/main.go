package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	// db.Ping()
	// models.CreateUser("Leo", "Leo123", "leoame@gmail.com")
	user := models.GetUserByID(1)
	fmt.Println("User:", user)
	user.Username = "LEO UPDATED"
	user.Save()
	fmt.Println("Updated User:", user)
	// db.CreateTable(models.UserSchema, "users")
	defer db.Close()
}
