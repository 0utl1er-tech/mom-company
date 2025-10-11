package contact

import (
	"context"

	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/google/uuid"
)

func (s *Service) CreateContact(ctx context.Context, req *contactv1.ContactRequest) (*contactv1.Contact, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	contact_id := uuid.New()

	// 連絡先を作成
	contact, err := s.db.CreateContact(ctx, db.CreateContactParams{
		ID:    contact_id,
		Email: req.Email,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &contactv1.Contact{
		Id:    contact.ID.String(),
		Email: contact.Email,
		Phone: contact.Phone,
	}, nil
}
