package postgres

import (
	"context"

	"github.com/chnejohnson/shortener/domain"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type pgAccountRepo struct {
	conn *pgx.Conn
}

// NewRepository ...
func NewRepository(conn *pgx.Conn) domain.AccountRepository {
	return &pgAccountRepo{conn}
}

// Create ...
func (p *pgAccountRepo) Create(a *domain.Account) error {
	sql := `
		INSERT INTO public.users (name, email, password)
		VALUES($1, $2, $3);`

	_, err := p.conn.Exec(context.Background(), sql, a.Name, a.Email, a.Password)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.WithFields(logrus.Fields{
		"name":     a.Name,
		"email":    a.Email,
		"password": a.Password,
	}).Info("Succeed to store account into database")

	return nil
}

// Fetch ...
func (p *pgAccountRepo) Find(acc string) (string, error) {
	sql := `
	SELECT password
	FROM users 
	WHERE email = $1;`

	var password string
	err := p.conn.QueryRow(context.Background(), sql, acc).Scan(&password)
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	return password, nil
}

// Update ...
// func (p *pgAccountRepo) Update(account *domain.Account) error {

// }

// Delete ...
// func (p *pgAccountRepo) Delete(account *domain.Account) error {

// }
