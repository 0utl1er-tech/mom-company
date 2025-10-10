package service

import db "github.com/0utl1er-tech/mom-company/gen/sqlc"

type CompanyService struct {
	db *db.Queries
}

func NewCompanyService(db *db.Queries) *CompanyService {
	return &CompanyService{db: db}
}
