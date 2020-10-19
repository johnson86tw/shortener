package service

import (
	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
)

type UserURLService struct {
	db domain.UserURLRepository
}

func NewUserURLService(db domain.UserURLRepository) domain.UserURLService {
	return &UserURLService{db}
}

func (u *UserURLService) FetchAll(userID uuid.UUID) ([]*domain.UserURL, error) {
	urls, err := u.db.FetchAll(userID)
	if err != nil {
		return urls, err
	}

	return urls, nil
}
