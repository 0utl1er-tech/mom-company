package contact

import (
	"context"

	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/google/uuid"
)

func (s *Service) UpdateContact(ctx context.Context, id string, req *contactv1.ContactRequest) (*contactv1.Contact, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	contactID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// 連絡先を更新
	contact, err := s.db.UpdateContact(ctx, db.UpdateContactParams{
		ID:    contactID,
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
