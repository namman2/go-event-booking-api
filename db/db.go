package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DBase *sql.DB

func InitDB() {
	var err error
	DBase, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Failed to initialize the SQLite database")
	}

	// Pool of ongoing open Connections
	DBase.SetMaxOpenConns(10)
	DBase.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER REFERENCES users(id)
	)
	`
	_, err := DBase.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}
	_, err = DBase.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}

}
