package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	username string
}

// Fetch a user by username from the database
func FetchUser(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	u := User{}
	err := row.Scan(&u)

	return &u, err
}

// Add a user to the database, hashing the provided raw password value
func AddUser(db *sql.DB, username string, passwordRaw string) error {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(passwordRaw), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, string(passwordHashed))

	return err
}
