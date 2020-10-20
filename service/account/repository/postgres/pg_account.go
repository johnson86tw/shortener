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
		return err
	}

	logrus.WithFields(logrus.Fields{
		"name":     a.Name,
		"email":    a.Email,
		"password": a.Password,
	}).Info("Succeed to store account into database")

	return nil
}

// Find ...
func (p *pgAccountRepo) Find(email string) (*domain.Account, error) {
	sql := `
	SELECT password, user_id
	FROM users 
	WHERE email = $1;`

	acc := &domain.Account{}

	rows, err := p.conn.Query(context.Background(), sql, email)
	if err != nil {
		return acc, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&acc.Password, &acc.UserID)
		if err != nil {
			return acc, err
		}
	}

	return acc, nil
}

// Update ...
// func (p *pgAccountRepo) Update(account *domain.Account) error {

// }

// Delete ...
// func (p *pgAccountRepo) Delete(account *domain.Account) error {

// }
