package mocks

import (
	"github.com/chnejohnson/shortener/domain"
	"github.com/stretchr/testify/mock"
)

type RedirectRepository struct {
	mock.Mock
}

func (m *RedirectRepository) Find(code string) (*domain.Redirect, error) {
	args := m.Called(code)

	return args.Get(0).(*domain.Redirect), args.Error(1)
}

func (m *RedirectRepository) Store(redirect *domain.Redirect) error {
	args := m.Called(redirect)

	var r0 error
	if rf, ok := args.Get(0).(func(*domain.Redirect) error); ok {
		r0 = rf(redirect)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (m *RedirectRepository) FindByURL(url string) (*domain.Redirect, error) {
	args := m.Called(url)

	return args.Get(0).(*domain.Redirect), args.Error(1)
}
