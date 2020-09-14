package domain

// Redirect ...
type Redirect struct {
	URL       string
	Code      string
	CreatedAt int64
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
