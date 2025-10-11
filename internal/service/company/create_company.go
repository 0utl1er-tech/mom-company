package company

import (
	"context"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateCompany(ctx context.Context, req *companyv1.CreateCompanyRequest) (*companyv1.CreateCompanyResponse, error) {
	company_id := uuid.New()
	company_contact_id := uuid.New()
	ceo_staff_id := uuid.New()
	ceo_contact_id := uuid.New()

	// トランザクションを使用して循環参照の問題を解決
	tx, err := s.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// トランザクション用のクエリオブジェクトを作成
	txQueries := s.db.WithTx(tx)

	// 一時的に外部キー制約を無効化
	_, err = tx.Exec(ctx, "SET CONSTRAINTS ALL DEFERRED")
	if err != nil {
		return nil, err
	}

	// まず会社の連絡先を作成
	company_contact, err := txQueries.CreateContact(ctx, db.CreateContactParams{
		ID:    company_contact_id,
		Email: req.Contact.Email,
		Phone: req.Contact.Phone,
	})
	if err != nil {
		return nil, err
	}

	// CEOの連絡先を作成
	ceo_contact, err := txQueries.CreateContact(ctx, db.CreateContactParams{
		ID:    ceo_contact_id,
		Email: req.Ceo.Contact.Email,
		Phone: req.Ceo.Contact.Phone,
	})
	if err != nil {
		return nil, err
	}

	// 会社を作成（CEOは一時的にNULL）
	company, err := txQueries.CreateCompany(ctx, db.CreateCompanyParams{
		ID:          company_id,
		Ceo:         pgtype.UUID{Valid: false}, // 一時的にNULL
		Trademark:   req.Trademark,
		Type:        db.Type(req.Type),
		Position:    db.Presuf(req.Position),
		Address:     req.Address,
		CompanyCode: req.CompanyCode,
		ContactID:   company_contact_id,
	})
	if err != nil {
		return nil, err
	}

	// CEOスタッフを作成
	ceo_staff, err := txQueries.CreateStaff(ctx, db.CreateStaffParams{
		ID:        ceo_staff_id,
		Name:      req.Ceo.Name,
		Role:      req.Ceo.Role,
		ContactID: ceo_contact_id,
		CompanyID: company_id,
	})
	if err != nil {
		return nil, err
	}

	// CEOを更新
	_, err = txQueries.UpdateCompanyCeo(ctx, db.UpdateCompanyCeoParams{
		ID:  company_id,
		Ceo: pgtype.UUID{Bytes: ceo_staff_id, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// トランザクションをコミット
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	// レスポンス用のContactを作成
	contact := &contactv1.Contact{
		Id:    company_contact.ID.String(),
		Email: company_contact.Email,
		Phone: company_contact.Phone,
	}

	// レスポンス用のCEO Staffを作成
	ceo := &staffv1.Staff{
		Id:   ceo_staff.ID.String(),
		Name: ceo_staff.Name,
		Role: ceo_staff.Role,
		Contact: &contactv1.Contact{
			Id:    ceo_contact.ID.String(),
			Email: ceo_contact.Email,
			Phone: ceo_contact.Phone,
		},
	}

	return &companyv1.CreateCompanyResponse{
		Company: &companyv1.Company{
			Id:          company.ID.String(),
			Trademark:   company.Trademark,
			Type:        string(company.Type),
			Position:    string(company.Position),
			Address:     company.Address,
			CompanyCode: company.CompanyCode,
			Contact:     contact,
			Staff:       []*staffv1.Staff{ceo},
			CreatedAt:   timestamppb.New(company.CreatedAt),
		},
	}, nil
}
