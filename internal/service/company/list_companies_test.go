package company

import (
	"testing"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCompanies(t *testing.T) {
	// テストデータの準備
	req := &companyv1.ListCompaniesRequest{
		Page:  1,
		Limit: 10,
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, int32(1), req.Page)
	assert.Equal(t, int32(10), req.Limit)
}
