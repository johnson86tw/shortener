package domain

import (
	"time"

	"github.com/google/uuid"
)

// Redirect ...
type Redirect struct {
	URL        string    `json:"url"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	TotalClick int       `json:"-"`
	UserID     uuid.UUID `json:"-"`
}

// RedirectService ...
type RedirectService interface {
	Redirect(code string) (string, error)
	Store(*Redirect) error
}

// RedirectRepository ...
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(*Redirect) error
	FindByURL(url string) (*Redirect, error)
}
