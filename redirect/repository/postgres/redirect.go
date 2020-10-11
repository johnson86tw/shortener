package postgres

import (
	"context"

	"github.com/chnejohnson/shortener/domain"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type pgRedirectRepository struct {
	conn *pgx.Conn
}

// NewRepository ...
func NewRepository(conn *pgx.Conn) domain.RedirectRepository {
	return &pgRedirectRepository{conn}
}

func (p *pgRedirectRepository) Find(code string) (*domain.Redirect, error) {
	sql := "SELECT url, created_at FROM urls WHERE code=$1;"

	redirect := &domain.Redirect{}
	redirect.Code = code
	err := p.conn.QueryRow(context.Background(), sql, code).Scan(&redirect.URL, &redirect.CreatedAt)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.WithFields(logrus.Fields{"url": redirect.URL, "created_at": redirect.CreatedAt}).Info("Succeed to find redirect")
	return redirect, nil
}

func (p *pgRedirectRepository) Store(redirect *domain.Redirect) error {
	sql := "INSERT INTO public.urls (url, code) VALUES($1, $2);"

	_, err := p.conn.Exec(context.Background(), sql, redirect.URL, redirect.Code)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Succeed to store redirect into database")
	return nil
}
