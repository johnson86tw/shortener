package service

import (
	"errors"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AccountService ...
type AccountService struct {
	db domain.AccountRepository
}

// NewAccountService ...
func NewAccountService(db domain.AccountRepository) domain.AccountService {
	return &AccountService{db}
}

// Create ...
func (a *AccountService) Create(acc *domain.Account) error {
	if acc.Name == "" || acc.Email == "" || acc.Password == "" {
		return errors.New("Name, Email, or Password cannot be empty")
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	acc.Password = string(hash)

	err = a.db.Create(acc)
	if err != nil {
		return err
	}

	return nil
}

// Login ...
func (a *AccountService) Login(email string, password string) (uuid.UUID, error) {
	acc, err := a.db.Find(email)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		return uuid.UUID{}, err
	}

	return acc.UserID, nil
}
