package company

import (
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	db       *db.Queries
	connPool *pgxpool.Pool
}

func NewService(queries *db.Queries, connPool *pgxpool.Pool) *Service {
	return &Service{db: queries, connPool: connPool}
}
