package db

import "time"

// User model for authentication.
//
// swagger:model
type User struct {
	ID uint64 `json:"id" db:"id"`
	// User's username
	// required: true
	UserName string `json:"username" db:"username"`
	// User's password
	// required: true
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// GenericError error model.
//
// swagger:model
type GenericError struct {
	// Error massage.
	Error string `json:"error"`
}
