package domain

import (
	"time"
)

// Account ...
type Account struct {
	UserID    string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AccountService ...
type AccountService interface {
	Create(*Account) error
	Login(email string, password string) error
	// Update(*Account) error
	// Delete(*Account) error
}

// AccountRepository ...
type AccountRepository interface {
	Create(*Account) error
	Find(email string) (string, error)
	// Fetch(userID string) (*Account, error)
	// Update(*Account) error
	// Delete(*Account) error
}
