package domain

import "time"

// Redirect ...
type Redirect struct {
	URL       string
	Code      string
	CreatedAt time.Time
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
