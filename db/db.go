package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")

	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		log.Printf("Could not create users table: %v", err)
		panic("Could not create users table")
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users (id)
		)`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		log.Printf("Could not create events table: %v", err)
		panic("Could not create events table")
	}

	createRegistrationTable := `CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		eventId INTEGER,
		userId INTEGER,
		FOREIGN KEY (eventId) REFERENCES events (id),
		FOREIGN KEY (userId) REFERENCES users (id)
	)`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		log.Printf("Could not create registrations table: %v", err)
		panic("Could not create registrations table")
	}
}
