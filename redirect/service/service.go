package service

import (
	"time"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
)

// RedirectService ...
type redirectService struct {
	redirectRepo domain.RedirectRepository
}

// NewRedirectService ...
func NewRedirectService(repo domain.RedirectRepository) domain.RedirectService {
	return &redirectService{repo}
}

// Find ...
func (r *redirectService) Find(code string) (*domain.Redirect, error) {
	rdrt, err := r.redirectRepo.Find(code)
	if err != nil {
		return nil, err
	}

	return rdrt, nil
}

// Store ...
func (r *redirectService) Store(rdrt *domain.Redirect) error {
	u, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	rdrt.Code = u.String()
	rdrt.CreatedAt = time.Now().UTC().Unix()

	err = r.redirectRepo.Store(rdrt)
	if err != nil {
		return err
	}

	return nil
}
