package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not connect database")
	}
	DB.SetMaxOpenConns(10) // limiting the number of db connections we can have
	DB.SetMaxIdleConns(5)  // no. of connections we want to be open at all times
	createTable()
}

func createTable() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("could not create users table!!!!")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)

	)
	`
	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("could not create events table!!!!")
	}

}
