package mocks

import (
	"github.com/chnejohnson/shortener/domain"
	"github.com/stretchr/testify/mock"
)

type RedirectService struct {
	mock.Mock
}

func (m *RedirectService) Redirect(code string) (*domain.Redirect, error) {
	args := m.Called(code)

	return args.Get(0).(*domain.Redirect), args.Error(1)
}
