package company

import (
	"context"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) ListCompanies(ctx context.Context, req *companyv1.ListCompaniesRequest) (*companyv1.ListCompaniesResponse, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	companies, err := s.db.ListCompanies(ctx)
	if err != nil {
		return nil, err
	}

	var companyList []*companyv1.Company
	for _, company := range companies {
		// レスポンス用のContactを作成
		contact := &contactv1.Contact{
			Id:    company.ContactID.String(),
			Email: company.ContactEmail.String,
			Phone: company.ContactPhone.String,
		}

		companyList = append(companyList, &companyv1.Company{
			Id:          company.ID.String(),
			Trademark:   company.Trademark,
			Type:        string(company.Type),
			Position:    string(company.Position),
			Address:     company.Address,
			CompanyCode: company.CompanyCode,
			Contact:     contact,
			CreatedAt:   timestamppb.New(company.CreatedAt),
		})
	}

	return &companyv1.ListCompaniesResponse{
		Companies: companyList,
	}, nil
}
