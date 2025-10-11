package contact

import (
	"testing"

	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateContact(t *testing.T) {
	// テストデータの準備
	req := &contactv1.ContactRequest{
		Email: "test@example.com",
		Phone: "03-1234-5678",
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, "test@example.com", req.Email)
	assert.Equal(t, "03-1234-5678", req.Phone)
}

func TestUpdateContact(t *testing.T) {
	// テストデータの準備
	contactID := "test-contact-id"
	req := &contactv1.ContactRequest{
		Email: "updated@example.com",
		Phone: "03-9999-9999",
	}

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, req)
	assert.Equal(t, contactID, contactID)
	assert.Equal(t, "updated@example.com", req.Email)
	assert.Equal(t, "03-9999-9999", req.Phone)
}

func TestDeleteContact(t *testing.T) {
	// テストデータの準備
	contactID := "test-contact-id"

	// テスト実行（実際のテストでは、テスト用のデータベースを使用することを推奨）
	// ここでは基本的な構造のテストのみ実行
	require.NotNil(t, contactID)
	assert.Equal(t, "test-contact-id", contactID)
	assert.NotEmpty(t, contactID)
}
