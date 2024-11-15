package db

import (
	"database/sql"
	_ "log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Pastikan DSN sudah sesuai dengan koneksi ke MySQL Anda
	dsn := "root:@tcp(127.0.0.1:3306)/events_db?parseTime=true&loc=Local"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table: " + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table: " + err.Error())
	}

	// log.Println("Events table created or already exists")

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table: " + err.Error())
	}
}
