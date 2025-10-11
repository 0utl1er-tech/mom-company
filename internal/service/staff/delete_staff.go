package staff

import (
	"context"

	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/google/uuid"
)

func (s *Service) DeleteStaff(ctx context.Context, req *staffv1.DeleteStaffRequest) (*staffv1.DeleteStaffResponse, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	staffID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = s.db.DeleteStaff(ctx, staffID)
	if err != nil {
		return nil, err
	}

	return &staffv1.DeleteStaffResponse{
		Id: req.Id,
	}, nil
}
