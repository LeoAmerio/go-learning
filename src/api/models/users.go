package models

import "apirest/db"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const UserSchema string = `
CREATE TABLE IF NOT EXISTS users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(100) NOT NULL UNIQUE
)`

type Users []User

func NewUser(username, password, email string) *User {
	return &User{
		Username: username,
		Password: password,
		Email:    email,
	}
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

// Insertar registro de usuario
func (user *User) insert() {
	sql := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	res, _ := db.Exec(sql, user.Username, user.Password, user.Email)

	user.ID, _ = res.LastInsertId()
}

// ListAll registers
func ListUsers() (Users, error) {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int64) (*User, error) {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id = ?"
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	}

	return user, nil
}

// Update register
func (user *User) update() {
	sql := "UPDATE users SET username = ?, password = ?, email = ? WHERE id = ?"
	_, err := db.Exec(sql, user.Username, user.Password, user.Email, user.ID)
	if err != nil {
		panic(err.Error())
	}
}

// Save or update user
func (user *User) Save() {
	if user.ID == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(sql, user.ID)
	if err != nil {
		panic(err.Error())
	}
	user.ID = 0 // Reset ID after deletion
}
