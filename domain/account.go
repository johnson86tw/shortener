package domain

import (
	"time"

	"github.com/google/uuid"
)

// Account ...
type Account struct {
	UserID    uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AccountService ...
type AccountService interface {
	Create(*Account) error
	Login(email string, password string) (uuid.UUID, error)
	// Update(*Account) error
	// Delete(*Account) error
}

// AccountRepository ...
type AccountRepository interface {
	Create(*Account) error
	Find(email string) (*Account, error)
	// Fetch(userID string) (*Account, error)
	// Update(*Account) error
	// Delete(*Account) error
}
