package service_test

import (
	"errors"
	"testing"

	"github.com/chnejohnson/shortener/domain"
	"github.com/chnejohnson/shortener/domain/mocks"
	redirectService "github.com/chnejohnson/shortener/service/redirect/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert" // 失敗仍繼續
	"github.com/stretchr/testify/mock"
	// 失敗就停
)

func TestRedirect(t *testing.T) {

	t.Run("success without user id", func(t *testing.T) {
		mockRedirectRepository := new(mocks.RedirectRepository)
		mockUserURLRepository := new(mocks.UserURLRepository)
		mockRedirect := domain.Redirect{
			URL: "https://www.google.com",
		}
		mockRedirectRepository.On("Find", mock.AnythingOfType("string")).Return(&mockRedirect, nil).Once()
		mockUserURLRepository.On("AddTotalClick", mock.AnythingOfType("string")).Return(nil).Once()

		rs := redirectService.NewRedirectService(mockRedirectRepository, mockUserURLRepository)

		redirect, err := rs.Redirect("123456")
		assert.NoError(t, err)
		assert.NotEmpty(t, redirect.URL)

		mockRedirectRepository.AssertExpectations(t)
		// 取出的 redirect 沒有 uuid，不可發動 AddTotalClick
		mockUserURLRepository.AssertNotCalled(t, "AddTotalClick", mock.AnythingOfType("string"))
	})

	t.Run("success with user id", func(t *testing.T) {
		mockRedirectRepository := new(mocks.RedirectRepository)
		mockUserURLRepository := new(mocks.UserURLRepository)
		mockRedirect := domain.Redirect{
			URL:    "https://www.google.com",
			UserID: uuid.New(),
		}
		mockRedirectRepository.On("Find", mock.AnythingOfType("string")).Return(&mockRedirect, nil).Once()
		mockUserURLRepository.On("AddTotalClick", mock.AnythingOfType("string")).Return(nil).Once()

		rs := redirectService.NewRedirectService(mockRedirectRepository, mockUserURLRepository)

		redirect, err := rs.Redirect("123456")
		assert.NoError(t, err)
		assert.NotEmpty(t, redirect.URL)

		mockRedirectRepository.AssertExpectations(t)
		// 取出的 redirect 有 uuid，必須發動 AddTotalClick
		mockUserURLRepository.AssertExpectations(t)
	})

	t.Run("error happens in finding url", func(t *testing.T) {
		mockRedirectRepository := new(mocks.RedirectRepository)
		mockUserURLRepository := new(mocks.UserURLRepository)
		mockRedirect := domain.Redirect{
			URL: "https://www.google.com",
		}

		mockRedirectRepository.On("Find", mock.AnythingOfType("string")).Return(&mockRedirect, errors.New("error")).Once()

		rs := redirectService.NewRedirectService(mockRedirectRepository, mockUserURLRepository)

		redirect, err := rs.Redirect("123456")
		assert.Error(t, err)
		assert.Empty(t, redirect)

		mockRedirectRepository.AssertExpectations(t)
	})

	t.Run("error happens in adding total click", func(t *testing.T) {
		mockRedirectRepository := new(mocks.RedirectRepository)
		mockUserURLRepository := new(mocks.UserURLRepository)
		mockRedirect := domain.Redirect{
			URL:    "https://www.google.com",
			UserID: uuid.New(),
		}
		mockRedirectRepository.On("Find", mock.AnythingOfType("string")).Return(&mockRedirect, nil).Once()
		mockUserURLRepository.On("AddTotalClick", mock.AnythingOfType("string")).Return(errors.New("error")).Once()

		rs := redirectService.NewRedirectService(mockRedirectRepository, mockUserURLRepository)

		redirect, err := rs.Redirect("123456")
		assert.Error(t, err)
		assert.Equal(t, &domain.Redirect{}, redirect)

		mockRedirectRepository.AssertExpectations(t)
		mockUserURLRepository.AssertExpectations(t)
	})

}
