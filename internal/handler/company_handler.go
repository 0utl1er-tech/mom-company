package handler

import (
	"context"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	"github.com/0utl1er-tech/mom-company/internal/service/company"
	"github.com/bufbuild/connect-go"
)

// CompanyServiceHandler はConnectのハンドラーインターフェースを実装します
type CompanyServiceHandler struct {
	*company.Service
}

// NewCompanyServiceHandler は新しいCompanyServiceHandlerを作成します
func NewCompanyServiceHandler(service *company.Service) *CompanyServiceHandler {
	return &CompanyServiceHandler{Service: service}
}

// CreateCompany は会社を作成します
func (h *CompanyServiceHandler) CreateCompany(ctx context.Context, req *connect.Request[companyv1.CreateCompanyRequest]) (*connect.Response[companyv1.CreateCompanyResponse], error) {
	resp, err := h.Service.CreateCompany(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// GetCompany は会社を取得します
func (h *CompanyServiceHandler) GetCompany(ctx context.Context, req *connect.Request[companyv1.GetCompanyRequest]) (*connect.Response[companyv1.GetCompanyResponse], error) {
	resp, err := h.Service.GetCompany(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// ListCompanies は会社一覧を取得します
func (h *CompanyServiceHandler) ListCompanies(ctx context.Context, req *connect.Request[companyv1.ListCompaniesRequest]) (*connect.Response[companyv1.ListCompaniesResponse], error) {
	resp, err := h.Service.ListCompanies(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
