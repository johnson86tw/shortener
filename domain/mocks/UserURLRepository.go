package mocks

import (
	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserURLRepository struct {
	mock.Mock
}

func (m *UserURLRepository) Find(code string) (string, error) {
	args := m.Called(code)
	return args.String(0), args.Error(1)
}

func (m *UserURLRepository) FetchAll(userID uuid.UUID) ([]*domain.UserURL, error) {
	args := m.Called(userID)
	return args.Get(0).([]*domain.UserURL), args.Error(1)
}

func (m *UserURLRepository) AddTotalClick(code string) error {
	args := m.Called(code)
	return args.Error(0)
}

func (m *UserURLRepository) AddURL(userURL *domain.UserURL) error {
	args := m.Called(userURL)
	return args.Error(0)
}
