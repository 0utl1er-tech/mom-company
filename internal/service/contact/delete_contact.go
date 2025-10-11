package contact

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) DeleteContact(ctx context.Context, id string) error {
	contactID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.db.DeleteContact(ctx, contactID)
}
