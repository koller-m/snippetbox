package models

import (
	"database/sql"
	"time"
)

// Define User type
// Fields align with users table in db
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Define UserModel type which wraps db connection pool
type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authentication method returns user ID if valid email and password
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Check if the users exists with the specific ID
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
