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
	// user := models.GetUserByID(2)
	// fmt.Println("User:", user)
	// user.Username = "LEO UPDATED"
	// user.Delete()
	// fmt.Println("Updated User:", user)
	db.TruncateTable("users")
	fmt.Println(models.ListUsers())
	// user.Delete(1)
	defer db.Close()
}
