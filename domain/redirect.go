package domain

import (
	"time"
)

// Redirect ...
type Redirect struct {
	URL         string    `json:"url"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time
	TotalClicks int
	Owner       int
}

// RedirectService ...
type RedirectService interface {
	Find(string) (*Redirect, error)
	Store(*Redirect) error
}

// RedirectRepository ...
type RedirectRepository interface {
	Find(string) (*Redirect, error)
	Store(*Redirect) error
}

// UserRedirectService ...
type UserRedirectService interface {
	FetchAll(userID string) ([]*Redirect, error)
	Add(userID string, redirect *Redirect) error
	Edit(code string) error
	Delete(code string) error
}

// UserRedirectRepository ...
type UserRedirectRepository interface {
	FetchAll(userID string) ([]*Redirect, error)
	Add(userID string, redirect *Redirect) error
	Edit(code string) error
	Delete(code string) error
}
