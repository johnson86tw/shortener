package domain

import "time"

// Redirect ...
type Redirect struct {
	URL       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"createdAt"`
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
