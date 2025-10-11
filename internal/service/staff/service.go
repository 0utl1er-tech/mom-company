package staff

import db "github.com/0utl1er-tech/mom-company/gen/sqlc"

type Service struct {
	db *db.Queries
}

func NewService(db *db.Queries) *Service {
	return &Service{db: db}
}
