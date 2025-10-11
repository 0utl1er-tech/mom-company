package company

import (
	"context"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateCompany(ctx context.Context, req *companyv1.CreateCompanyRequest) (*companyv1.CreateCompanyResponse, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	company_id := uuid.New()
	company_contact_id := uuid.New()

	// トランザクションを使用
	tx, err := s.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// トランザクション用のクエリオブジェクトを作成
	txQueries := s.db.WithTx(tx)

	// 会社の連絡先を作成
	company_contact, err := txQueries.CreateContact(ctx, db.CreateContactParams{
		ID:    company_contact_id,
		Email: req.Contact.Email,
		Phone: req.Contact.Phone,
	})
	if err != nil {
		return nil, err
	}

	// 会社を作成
	company, err := txQueries.CreateCompany(ctx, db.CreateCompanyParams{
		ID:          company_id,
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

	// スタッフを作成（もしリクエストに含まれている場合）
	var staffList []*staffv1.Staff
	for _, staffReq := range req.Staff {
		staff_id := uuid.New()
		staff_contact_id := uuid.New()

		// スタッフの連絡先を作成
		staff_contact, err := txQueries.CreateContact(ctx, db.CreateContactParams{
			ID:    staff_contact_id,
			Email: staffReq.Contact.Email,
			Phone: staffReq.Contact.Phone,
		})
		if err != nil {
			return nil, err
		}

		// スタッフを作成
		staff, err := txQueries.CreateStaff(ctx, db.CreateStaffParams{
			ID:        staff_id,
			Name:      staffReq.Name,
			Role:      staffReq.Role,
			ContactID: staff_contact_id,
			CompanyID: company_id,
		})
		if err != nil {
			return nil, err
		}

		// レスポンス用のStaffを作成
		staffList = append(staffList, &staffv1.Staff{
			Id:   staff.ID.String(),
			Name: staff.Name,
			Role: staff.Role,
			Contact: &contactv1.Contact{
				Id:    staff_contact.ID.String(),
				Email: staff_contact.Email,
				Phone: staff_contact.Phone,
			},
		})
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

	return &companyv1.CreateCompanyResponse{
		Company: &companyv1.Company{
			Id:          company.ID.String(),
			Trademark:   company.Trademark,
			Type:        string(company.Type),
			Position:    string(company.Position),
			Address:     company.Address,
			CompanyCode: company.CompanyCode,
			Contact:     contact,
			Staff:       staffList,
			CreatedAt:   timestamppb.New(company.CreatedAt),
		},
	}, nil
}
