package ad

import "github.com/jackc/pgx/v5"

type AdRepository struct {
	db *pgx.Conn
}

func NewAdRepository(db *pgx.Conn) *AdRepository {
	return &AdRepository{db: db}
}
