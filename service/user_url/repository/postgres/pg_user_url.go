package postgres

import (
	"context"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type pgUserURLRepository struct {
	conn *pgx.Conn
}

// NewRepository ...
func NewRepository(conn *pgx.Conn) domain.UserURLRepository {
	return &pgUserURLRepository{conn}
}

func (p *pgUserURLRepository) Find(code string) (string, error) {
	sql := "SELECT url FROM user_urls WHERE code=$1;"

	url := ""
	err := p.conn.QueryRow(context.Background(), sql, code).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (p *pgUserURLRepository) FetchAll(userID uuid.UUID) ([]*domain.UserURL, error) {
	sql := `
		SELECT id, url, code, created_at, total_click
		FROM user_urls 
		WHERE user_id = $1;`

	urls := []*domain.UserURL{}

	rows, err := p.conn.Query(context.Background(), sql, userID.String())
	if err != nil {
		return urls, err
	}

	defer rows.Close()

	for rows.Next() {
		url := UserURL{}

		err := rows.Scan(&url.ID, &url.URL, &url.Code, &url.CreatedAt, &url.TotalClick)
		if err != nil {
			return urls, err
		}

		durl := &domain.UserURL{
			ID:         url.ID,
			URL:        url.URL,
			Code:       url.Code,
			CreatedAt:  url.CreatedAt.Time,
			TotalClick: url.TotalClick,
		}

		if err != nil {
			return urls, err
		}

		urls = append(urls, durl)
	}

	if rows.Err() != nil {
		return urls, rows.Err()
	}

	return urls, nil

}

// AddTotalClick ...
func (p *pgUserURLRepository) AddTotalClick(code string) error {
	sql := `
		UPDATE user_urls
		SET total_click = total_click + 1
		WHERE code = $1;
	`
	_, err := p.conn.Exec(context.Background(), sql, code)
	if err != nil {
		return err
	}

	return nil

}

// AddURL ...
func (p *pgUserURLRepository) AddURL(uu *domain.UserURL) error {
	sql := `
		INSERT INTO user_urls (url, code, user_id) 
		VALUES($1, $2, $3);`

	_, err := p.conn.Exec(context.Background(), sql, uu.URL, uu.Code, uu.UserID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url":    uu.URL,
			"code":   uu.Code,
			"userID": uu.UserID,
		}).Error("Fail to insert data into postgres")
		return err
	}

	return nil
}
