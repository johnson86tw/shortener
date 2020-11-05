package postgres

import (
	"context"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
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
	sql := `
		SELECT u.url, total_click, user_id 
		FROM (SELECT url, code FROM urls union SELECT url, code FROM user_urls) AS u
		LEFT OUTER JOIN user_urls uu
		ON uu.code = u.code
		WHERE u.code = $1;`

	redirect := &domain.Redirect{}
	var totalClick *int
	var userID *uuid.UUID

	err := p.conn.QueryRow(context.Background(), sql, code).Scan(&redirect.URL, &totalClick, &userID)
	if err != nil {
		return nil, err
	}

	if userID != nil && totalClick != nil {
		redirect.TotalClick = *totalClick
		redirect.UserID = *userID
	}

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

func (p *pgRedirectRepository) FindByURL(url string) (*domain.Redirect, error) {
	sql := "SELECT code FROM urls WHERE url = $1"

	redirect := &domain.Redirect{}
	rows, err := p.conn.Query(context.Background(), sql, url)
	if err != nil {
		return redirect, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&redirect.Code)
		if err != nil {
			return redirect, err
		}

		// 只取第一個
		break
	}

	return redirect, nil

}
