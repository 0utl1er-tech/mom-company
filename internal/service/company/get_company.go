package company

import (
	"context"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	"github.com/0utl1er-tech/mom-company/internal/util"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetCompany(ctx context.Context, req *companyv1.GetCompanyRequest) (*companyv1.GetCompanyResponse, error) {
	// リクエストをバリデーション
	if err := util.ValidateMessage(req); err != nil {
		return nil, err
	}

	companyID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	company, err := s.db.GetCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	staffs, err := s.db.ListStaff(ctx, companyID)
	if err != nil {
		return nil, err
	}

	// レスポンス用のContactを作成
	contact := &contactv1.Contact{
		Id:    company.ContactID.String(),
		Email: company.ContactEmail.String,
		Phone: company.ContactPhone.String,
	}

	staffList := []*staffv1.Staff{}
	for _, staff := range staffs {
		staffList = append(staffList, &staffv1.Staff{
			Id:   staff.ID.String(),
			Name: staff.Name,
			Role: staff.Role,
			Contact: &contactv1.Contact{
				Id:    staff.ContactID.String(),
				Email: staff.ContactEmail.String,
				Phone: staff.ContactPhone.String,
			},
		})
	}

	return &companyv1.GetCompanyResponse{
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
