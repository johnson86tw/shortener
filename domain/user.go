package domain

import "time"

// User ...
type User struct {
	UserID    string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserService ...
type UserService interface {
	Create(*User) error
	Fetch(userID string) *User
	Update(*User) error
	Delete(*User) error
}

// UserRepository ...
type UserRepository interface {
	Create(*User) error
	Fetch(userID string) *User
	Update(*User) error
	Delete(*User) error
}
