package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:sasa1234@tcp(localhost:3306)/goweb_db"

var db *sql.DB

func Connect() {
	sql, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connection established")
	db = sql
}

func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Database connection closed")
	} else {
		fmt.Println("No database connection to close")
	}
}

// VerifyConnection checks if the database connection is alive
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection is alive")
}

// Verify if the table exists
func TableExists(tableName string) bool {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	exists := rows.Next()
	return exists
}

// Create a table if it does not exist
func CreateTable(schema string, name string) {
	if !TableExists(name) {
		fmt.Println("Creating table if it does not exist...")
		res, err := Exec(schema)
		if err != nil {
			panic(err.Error())
		}
		rowsAffected, _ := res.RowsAffected()
		fmt.Println("Rows affected: ", rowsAffected)
	}
}

// Reset the database by truncate the table
func TruncateTable(tableName string) {
	if TableExists(tableName) {
		fmt.Println("Truncatting table...")
		query := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
		res, err := Exec(query)
		if err != nil {
			panic(err.Error())
		}
		rowsAffected, _ := res.RowsAffected()
		fmt.Println("Rows affected: ", rowsAffected)
	} else {
		fmt.Println("Table does not exist, no action taken.")
	}
}

// Polimorfismo en Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	defer Close()
	res, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Polmorfism en Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	defer Close()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
