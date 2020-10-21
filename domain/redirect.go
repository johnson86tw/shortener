package domain

import (
	"time"

	"github.com/google/uuid"
)

// Redirect ...
type Redirect struct {
	URL        string    `json:"url"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time
	TotalClick int
	UserID     uuid.UUID
}

// RedirectService ...
type RedirectService interface {
	Redirect(code string) (*Redirect, error)
	Store(*Redirect) error
}

// RedirectRepository ...
type RedirectRepository interface {
	Find(string) (*Redirect, error)
	Store(*Redirect) error
}
