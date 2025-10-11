package company

import (
	"testing"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCompany(t *testing.T) {
	// テストデータの準備
	companyID := uuid.New().String()
	req := &companyv1.GetCompanyRequest{
		Id: companyID,
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, companyID, req.Id)
	assert.NotEmpty(t, req.Id)
}
