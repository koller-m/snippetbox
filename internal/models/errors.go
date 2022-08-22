package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")

	// Error if a user tries to login with an incorrect email or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// Error if duplicate email is used
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
