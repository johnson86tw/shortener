package postgres

import (
	"github.com/jackc/pgtype"
)

type UserURL struct {
	ID         int
	URL        string
	Code       string
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
	DeletedAt  pgtype.Timestamp
	TotalClick int
	UserID     pgtype.UUID
}
