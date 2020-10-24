package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserURL struct {
	ID         int
	URL        string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	TotalClick int
	UserID     uuid.UUID
}

// UserURLService ...
type UserURLService interface {
	FetchAll(userID uuid.UUID) ([]*UserURL, error)
	AddURL(*UserURL) error
	// Edit(code string) error
	// Delete(code string) error
}

// UserURLRepository ...
type UserURLRepository interface {
	Find(code string) (string, error)
	FetchAll(userID uuid.UUID) ([]*UserURL, error)
	AddTotalClick(code string) error
	AddURL(*UserURL) error
	// Update(*UserURL) error
	// Delete(code string) error
}
