package service

import (
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

// Find ...
func (r *RedirectService) Redirect(code string) (*domain.Redirect, error) {
	redirect, err := r.db.Find(code)
	if err != nil {
		return nil, err
	}

	u := new(uuid.UUID)

	// record totalclick
	if redirect.UserID != *u {
		r.userURLDB.AddTotalClick(code)
	}

	return redirect, nil
}

// Store ...
func (r *RedirectService) Store(redirect *domain.Redirect) error {
	if redirect.URL == "" {
		logrus.Error("Redirect URL should not be empty")
		return errors.New("Redirect URL should not be empty")
	}
	// make redirect code
	code := genURLCode()
	logrus.WithField("code", code).Info("URL code has been generated")

	redirect.Code = code

	err := r.db.Store(redirect)
	if err != nil {
		logrus.Error("Fail to store redirect")
		return err
	}

	return nil
}

func genURLCode() string {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	b := make([]byte, 4)
	r.Read(b)
	s := hex.EncodeToString(b)
	return s
}
