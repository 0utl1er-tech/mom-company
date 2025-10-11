package util

import (
	"testing"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
)

func TestValidator(t *testing.T) {
	validator, err := NewValidator()
	if err != nil {
		t.Fatalf("バリデーターの作成に失敗: %v", err)
	}

	// 正常なCreateCompanyRequestのテスト
	t.Run("正常なCreateCompanyRequest", func(t *testing.T) {
		validRequest := &companyv1.CreateCompanyRequest{
			Trademark:   "テスト会社",
			Type:        "株式会社",
			Position:    "東京都渋谷区",
			Address:     "東京都渋谷区1-1-1",
			CompanyCode: "1234567890123",
			Contact: &contactv1.ContactRequest{
				Email: "test@example.com",
				Phone: "03-1234-5678",
			},
		}

		if err := validator.Validate(validRequest); err != nil {
			t.Errorf("正常なリクエストでバリデーションエラー: %v", err)
		}
	})

	// 無効なCreateCompanyRequestのテスト（空のtrademark）
	t.Run("無効なCreateCompanyRequest_空のtrademark", func(t *testing.T) {
		invalidRequest := &companyv1.CreateCompanyRequest{
			Trademark:   "", // 空文字列
			Type:        "株式会社",
			Position:    "東京都渋谷区",
			Address:     "東京都渋谷区1-1-1",
			CompanyCode: "1234567890123",
			Contact: &contactv1.ContactRequest{
				Email: "test@example.com",
				Phone: "03-1234-5678",
			},
		}

		if err := validator.Validate(invalidRequest); err == nil {
			t.Error("無効なリクエストがバリデーションを通過してしまった")
		}
	})

	// 無効なemailのテスト
	t.Run("無効なemail", func(t *testing.T) {
		invalidEmailRequest := &companyv1.CreateCompanyRequest{
			Trademark:   "テスト会社",
			Type:        "株式会社",
			Position:    "東京都渋谷区",
			Address:     "東京都渋谷区1-1-1",
			CompanyCode: "1234567890123",
			Contact: &contactv1.ContactRequest{
				Email: "invalid-email", // 無効なemail
				Phone: "03-1234-5678",
			},
		}

		if err := validator.Validate(invalidEmailRequest); err == nil {
			t.Error("無効なemailがバリデーションを通過してしまった")
		}
	})

	// 無効なcompany_codeのテスト（桁数が足りない）
	t.Run("無効なcompany_code", func(t *testing.T) {
		invalidCodeRequest := &companyv1.CreateCompanyRequest{
			Trademark:   "テスト会社",
			Type:        "株式会社",
			Position:    "東京都渋谷区",
			Address:     "東京都渋谷区1-1-1",
			CompanyCode: "123456789012", // 12桁は無効（13桁必要）
			Contact: &contactv1.ContactRequest{
				Email: "test@example.com",
				Phone: "03-1234-5678",
			},
		}

		if err := validator.Validate(invalidCodeRequest); err == nil {
			t.Error("無効なcompany_code（桁数不足）がバリデーションを通過してしまった")
		}
	})

	// 正常なCreateStaffRequestのテスト
	t.Run("正常なCreateStaffRequest", func(t *testing.T) {
		validStaffRequest := &staffv1.CreateStaffRequest{
			Name:      "田中太郎",
			Role:      "エンジニア",
			CompanyId: "company-123",
			Contact: &contactv1.ContactRequest{
				Email: "tanaka@example.com",
				Phone: "090-1234-5678",
			},
		}

		if err := validator.Validate(validStaffRequest); err != nil {
			t.Errorf("正常なスタッフリクエストでバリデーションエラー: %v", err)
		}
	})

	// ListCompaniesRequestのテスト
	t.Run("正常なListCompaniesRequest", func(t *testing.T) {
		validListRequest := &companyv1.ListCompaniesRequest{
			Page:  1,
			Limit: 50,
		}

		if err := validator.Validate(validListRequest); err != nil {
			t.Errorf("正常なリストリクエストでバリデーションエラー: %v", err)
		}
	})

	// 無効なlimitのテスト
	t.Run("無効なlimit", func(t *testing.T) {
		invalidListRequest := &companyv1.ListCompaniesRequest{
			Page:  1,
			Limit: 150, // 上限を超える
		}

		if err := validator.Validate(invalidListRequest); err == nil {
			t.Error("無効なlimitがバリデーションを通過してしまった")
		}
	})
}

func TestGlobalValidator(t *testing.T) {
	// グローバルバリデーターのテスト
	validRequest := &companyv1.CreateCompanyRequest{
		Trademark:   "テスト会社",
		Type:        "株式会社",
		Position:    "東京都渋谷区",
		Address:     "東京都渋谷区1-1-1",
		CompanyCode: "1234567890123",
		Contact: &contactv1.ContactRequest{
			Email: "test@example.com",
			Phone: "03-1234-5678",
		},
	}

	if err := ValidateMessage(validRequest); err != nil {
		t.Errorf("グローバルバリデーターでエラー: %v", err)
	}
}
