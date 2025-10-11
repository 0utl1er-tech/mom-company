package handler

import (
	"context"

	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	"github.com/0utl1er-tech/mom-company/internal/service/staff"
	"github.com/bufbuild/connect-go"
)

// StaffServiceHandler はConnectのハンドラーインターフェースを実装します
type StaffServiceHandler struct {
	*staff.Service
}

// NewStaffServiceHandler は新しいStaffServiceHandlerを作成します
func NewStaffServiceHandler(service *staff.Service) *StaffServiceHandler {
	return &StaffServiceHandler{Service: service}
}

// CreateStaff はスタッフを作成します
func (h *StaffServiceHandler) CreateStaff(ctx context.Context, req *connect.Request[staffv1.CreateStaffRequest]) (*connect.Response[staffv1.CreateStaffResponse], error) {
	resp, err := h.Service.CreateStaff(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// UpdateStaff はスタッフを更新します
func (h *StaffServiceHandler) UpdateStaff(ctx context.Context, req *connect.Request[staffv1.UpdateStaffRequest]) (*connect.Response[staffv1.UpdateStaffResponse], error) {
	resp, err := h.Service.UpdateStaff(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// DeleteStaff はスタッフを削除します
func (h *StaffServiceHandler) DeleteStaff(ctx context.Context, req *connect.Request[staffv1.DeleteStaffRequest]) (*connect.Response[staffv1.DeleteStaffResponse], error) {
	resp, err := h.Service.DeleteStaff(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
