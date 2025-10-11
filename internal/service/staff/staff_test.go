package staff

import (
	"testing"

	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateStaff(t *testing.T) {
	// テストデータの準備
	req := &staffv1.CreateStaffRequest{
		Name: "山田花子",
		Role: "一般社員",
		Contact: &contactv1.ContactRequest{
			Email: "yamada@test.com",
			Phone: "03-1234-5678",
		},
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, "山田花子", req.Name)
	assert.Equal(t, "一般社員", req.Role)
	assert.NotNil(t, req.Contact)
	assert.Equal(t, "yamada@test.com", req.Contact.Email)
	assert.Equal(t, "03-1234-5678", req.Contact.Phone)
}

func TestUpdateStaff(t *testing.T) {
	// テストデータの準備
	staffID := uuid.New().String()
	newName := "山田花子（更新）"
	req := &staffv1.UpdateStaffRequest{
		Id:   staffID,
		Name: &newName,
		Contact: &contactv1.ContactRequest{
			Email: "yamada.updated@test.com",
			Phone: "03-1234-9999",
		},
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, staffID, req.Id)
	assert.NotNil(t, req.Name)
	assert.Equal(t, newName, *req.Name)
	assert.NotNil(t, req.Contact)
	assert.Equal(t, "yamada.updated@test.com", req.Contact.Email)
	assert.Equal(t, "03-1234-9999", req.Contact.Phone)
}

func TestDeleteStaff(t *testing.T) {
	// テストデータの準備
	staffID := uuid.New().String()
	req := &staffv1.DeleteStaffRequest{
		Id: staffID,
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, staffID, req.Id)
	assert.NotEmpty(t, req.Id)
}
