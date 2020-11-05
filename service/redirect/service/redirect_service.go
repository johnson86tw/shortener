package service

import (
	"errors"

	"github.com/chnejohnson/shortener/domain"
	"github.com/chnejohnson/shortener/utils"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// RedirectService ...
type RedirectService struct {
	db        domain.RedirectRepository
	userURLDB domain.UserURLRepository
}

// NewRedirectService ...
func NewRedirectService(db domain.RedirectRepository, userURLDB domain.UserURLRepository) domain.RedirectService {
	return &RedirectService{db, userURLDB}
}

// Redirect ...
func (r *RedirectService) Redirect(code string) (*domain.Redirect, error) {
	redirect, err := r.db.Find(code)
	if err != nil {
		return &domain.Redirect{}, err
	}

	u := new(uuid.UUID)

	// record totalclick
	if redirect.UserID != *u {
		if err := r.userURLDB.AddTotalClick(code); err != nil {
			return &domain.Redirect{}, err
		}

	}

	return redirect, nil
}

// Store store the general url into database, if url exists, assigning exist code to Redirect
func (r *RedirectService) Store(redirect *domain.Redirect) error {
	if redirect.URL == "" {
		log.Error("Redirect URL should not be empty")
		return errors.New("Redirect URL should not be empty")
	}

	existed, err := r.db.FindByURL(redirect.URL)
	if err != nil {
		log.Error("FindByURL: ", err)
		return err
	}

	// 判斷 existed.Code 如果有值，直接對 redirect 賦值
	if existed.Code != "" {
		log.WithField("code", existed.Code).Info("Having existed URL")
		redirect.Code = existed.Code
		return nil
	}

	// make redirect code
	code := utils.GenURLCode()
	log.WithField("code", code).Info("URL code has been generated")

	redirect.Code = code

	err = r.db.Store(redirect)
	if err != nil {
		log.Error("Fail to store redirect")
		return err
	}

	return nil
}
