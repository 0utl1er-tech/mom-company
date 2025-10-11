package staff

import (
	"context"

	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) UpdateStaff(ctx context.Context, req *staffv1.UpdateStaffRequest) (*staffv1.UpdateStaffResponse, error) {
	// 既存のスタッフを取得
	staffID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	staff, err := s.db.GetStaff(ctx, staffID)
	if err != nil {
		return nil, err
	}

	// 連絡先を更新
	if req.Contact != nil {
		_, err = s.db.UpdateContact(ctx, db.UpdateContactParams{
			ID:    staff.ContactID,
			Email: req.Contact.Email,
			Phone: req.Contact.Phone,
		})
		if err != nil {
			return nil, err
		}
	}

	// スタッフ情報を更新
	updateParams := db.UpdateStaffParams{
		ID:        staffID,
		Name:      staff.Name,
		Role:      staff.Role,
		ContactID: staff.ContactID,
	}

	if req.Name != nil {
		updateParams.Name = *req.Name
	}
	if req.Role != nil {
		updateParams.Role = *req.Role
	}

	updatedStaff, err := s.db.UpdateStaff(ctx, updateParams)
	if err != nil {
		return nil, err
	}

	// 更新された連絡先を取得
	contact, err := s.db.GetContact(ctx, updatedStaff.ContactID)
	if err != nil {
		return nil, err
	}

	// レスポンス用のContactを作成
	contactResponse := &contactv1.Contact{
		Id:    contact.ID.String(),
		Email: contact.Email,
		Phone: contact.Phone,
	}

	return &staffv1.UpdateStaffResponse{
		Staff: &staffv1.Staff{
			Id:        updatedStaff.ID.String(),
			Name:      updatedStaff.Name,
			Role:      updatedStaff.Role,
			Contact:   contactResponse,
			CreatedAt: timestamppb.New(updatedStaff.CreatedAt),
		},
	}, nil
}
