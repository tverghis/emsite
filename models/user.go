package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Password string
}

// Fetch a user by username from the database
func FetchUser(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)

	u := User{}
	err := row.Scan(&u.Id, &u.Username, &u.Password)

	return &u, err
}

// Add a user to the database, hashing the provided raw password value
func AddUser(db *sql.DB, username string, passwordRaw string) error {
	password, err := hashPassword(passwordRaw)

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, string(password))

	return err
}

// Authenticate a user using the provided credentials.
// Returns a *User if the user exists and the credentials match, otherwise returns an error.
func AuthenticateUser(db *sql.DB, username string, passwordRaw string) (*User, error) {
	user, err := FetchUser(db, username)

	if user == nil || err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordRaw)); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(passwordRaw string) (string, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(passwordRaw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(passwordHashed), nil
}
