package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	// db.Ping()
	db.CreateTable(models.UserSchema, "users")
	defer db.Close()
}
