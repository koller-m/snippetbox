package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
	// Create a bcrypt hash of the plain-text password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
    VALUES(?, ?, ?, UTC_TIMESTAMP())`

	// Use Exec() method to insert user details and hashed password
	// Into the users table
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		// Use errors.As() to check if the error is type *mysql.MySQLError
		// If so, the error will be assigned to mySQLError variable
		// Then check if the error relates to users_uc_email key
		// By checking if error equals 1062
		// If true, return ErrDuplicateEmail error
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}

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
