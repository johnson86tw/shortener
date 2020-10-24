package service

import (
	"errors"

	"github.com/chnejohnson/shortener/domain"
	"github.com/chnejohnson/shortener/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// UserURLService ...
type UserURLService struct {
	db domain.UserURLRepository
}

// NewUserURLService ...
func NewUserURLService(db domain.UserURLRepository) domain.UserURLService {
	return &UserURLService{db}
}

// FetchAll ...
func (u *UserURLService) FetchAll(userID uuid.UUID) ([]*domain.UserURL, error) {
	urls, err := u.db.FetchAll(userID)
	if err != nil {
		return urls, err
	}

	return urls, nil
}

// AddURL ...
func (u *UserURLService) AddURL(uu *domain.UserURL) error {
	if uu.URL == "" || uu.UserID == *new(uuid.UUID) {
		err := errors.New("UserURL or UserID cannot be empty")
		logrus.New().WithFields(logrus.Fields{
			"url":    uu.URL,
			"userID": uu.UserID.String(),
		}).Error(err)
		return err
	}
	// generate url code
	code := utils.GenURLCode()
	uu.Code = code

	err := u.db.AddURL(uu)
	if err != nil {
		return err
	}

	return nil
}
