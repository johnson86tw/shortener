package service

import (
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/chnejohnson/shortener/domain"
	"github.com/sirupsen/logrus"
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
func (r *redirectService) Store(redirect *domain.Redirect) error {
	// make redirect code
	code := genURLCode()
	logrus.WithField("code", code).Info("URL code has been generated")

	redirect.Code = code

	err := r.redirectRepo.Store(redirect)
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
