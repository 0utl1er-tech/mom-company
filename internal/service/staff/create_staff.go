package staff

import (
	"context"

	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateStaff(ctx context.Context, req *staffv1.CreateStaffRequest) (*staffv1.CreateStaffResponse, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	staff_id := uuid.New()
	contact_id := uuid.New()

	// company_idをパース
	companyID, err := uuid.Parse(req.CompanyId)
	if err != nil {
		return nil, err
	}

	// 連絡先を作成
	contact, err := s.db.CreateContact(ctx, db.CreateContactParams{
		ID:    contact_id,
		Email: req.Contact.Email,
		Phone: req.Contact.Phone,
	})
	if err != nil {
		return nil, err
	}

	// スタッフを作成
	staff, err := s.db.CreateStaff(ctx, db.CreateStaffParams{
		ID:        staff_id,
		Name:      req.Name,
		Role:      req.Role,
		ContactID: contact_id,
		CompanyID: companyID,
	})
	if err != nil {
		return nil, err
	}

	// レスポンス用のContactを作成
	contactResponse := &contactv1.Contact{
		Id:    contact.ID.String(),
		Email: contact.Email,
		Phone: contact.Phone,
	}

	return &staffv1.CreateStaffResponse{
		Staff: &staffv1.Staff{
			Id:        staff.ID.String(),
			Name:      staff.Name,
			Role:      staff.Role,
			Contact:   contactResponse,
			CreatedAt: timestamppb.New(staff.CreatedAt),
		},
	}, nil
}
