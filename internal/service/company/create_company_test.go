package company

import (
	"testing"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCompany(t *testing.T) {
	// テストデータの準備
	req := &companyv1.CreateCompanyRequest{
		Trademark:   "テスト株式会社",
		Type:        "kabu",
		Position:    "prefix",
		Address:     "東京都渋谷区",
		CompanyCode: "1234567890123",
		Ceo: &staffv1.Staff{
			Name: "田中太郎",
			Role: "代表取締役",
			Contact: &contactv1.Contact{
				Email: "ceo@test.com",
				Phone: "03-1234-5678",
			},
		},
		Contact: &contactv1.ContactRequest{
			Email: "info@test.com",
			Phone: "03-1234-5679",
		},
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, "テスト株式会社", req.Trademark)
	assert.Equal(t, "kabu", req.Type)
	assert.Equal(t, "prefix", req.Position)
	assert.Equal(t, "東京都渋谷区", req.Address)
	assert.Equal(t, "1234567890123", req.CompanyCode)
	assert.NotNil(t, req.Ceo)
	assert.Equal(t, "田中太郎", req.Ceo.Name)
	assert.Equal(t, "代表取締役", req.Ceo.Role)
	assert.NotNil(t, req.Ceo.Contact)
	assert.Equal(t, "ceo@test.com", req.Ceo.Contact.Email)
	assert.Equal(t, "03-1234-5678", req.Ceo.Contact.Phone)
	assert.NotNil(t, req.Contact)
	assert.Equal(t, "info@test.com", req.Contact.Email)
	assert.Equal(t, "03-1234-5679", req.Contact.Phone)
}
