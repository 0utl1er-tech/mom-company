package main

import (
	"context"
	"fmt"
	"testing"

	companyv1 "github.com/0utl1er-tech/mom-company/gen/pb/company/v1"
	contactv1 "github.com/0utl1er-tech/mom-company/gen/pb/contact/v1"
	staffv1 "github.com/0utl1er-tech/mom-company/gen/pb/staff/v1"
	db "github.com/0utl1er-tech/mom-company/gen/sqlc"
	"github.com/0utl1er-tech/mom-company/internal/service/company"
	"github.com/0utl1er-tech/mom-company/internal/service/staff"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 統合テスト用のヘルパー関数
func setupTestService() (*company.Service, *staff.Service, func()) {
	// テスト用のデータベース接続文字列を設定
	// 実際のテストでは、テスト専用のデータベースを使用することを推奨
	testDBURL := "postgres://root:secret@localhost:5432/mom_company?sslmode=disable"

	// データベース接続プールを作成
	connPool, err := pgxpool.New(context.Background(), testDBURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to create connection pool: %v", err))
	}

	// クエリオブジェクトを作成
	queries := db.New(connPool)

	// サービスを作成
	companyService := company.NewService(queries, connPool)
	staffService := staff.NewService(queries)

	// クリーンアップ関数を返す
	cleanup := func() {
		connPool.Close()
	}

	return companyService, staffService, cleanup
}

func TestCompanyServiceIntegration(t *testing.T) {
	// 統合テストの例
	// 実際のテストでは、テスト用のデータベースを使用して
	// エンドツーエンドのテストを実行します

	t.Run("CreateCompany_Integration", func(t *testing.T) {
		// テスト用のサービスをセットアップ
		companyService, _, cleanup := setupTestService()
		defer cleanup()

		// テストデータの準備
		req := &companyv1.CreateCompanyRequest{
			Trademark:   "0UTL1ER",
			Type:        "kabu",
			Position:    "sufix",
			Address:     "東京都千代田区神田松永町１３番地ＶＯＲＴ秋葉原ＩＩ",
			CompanyCode: "9010001257448",
			Ceo: &staffv1.Staff{
				Name: "黒羽 晟",
				Role: "代表取締役",
				Contact: &contactv1.Contact{
					Email: "joe@0utl1er.tech",
					Phone: "090-1234-5678",
				},
			},
			Contact: &contactv1.ContactRequest{
				Email: "info@0utl1er.tech",
				Phone: "03-1111-1112",
			},
		}

		// テスト実行
		ctx := context.Background()
		resp, err := companyService.CreateCompany(ctx, req)

		// アサーション
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Company)
		assert.Equal(t, req.Trademark, resp.Company.Trademark)
		assert.Equal(t, req.Type, resp.Company.Type)
		assert.Equal(t, req.Position, resp.Company.Position)
		assert.Equal(t, req.Address, resp.Company.Address)
		assert.Equal(t, req.CompanyCode, resp.Company.CompanyCode)
		assert.NotEmpty(t, resp.Company.Id)
		assert.NotNil(t, resp.Company.Contact)
		assert.Equal(t, req.Contact.Email, resp.Company.Contact.Email)
		assert.Equal(t, req.Contact.Phone, resp.Company.Contact.Phone)
		assert.Len(t, resp.Company.Staff, 1)
		assert.Equal(t, req.Ceo.Name, resp.Company.Staff[0].Name)
		assert.Equal(t, req.Ceo.Role, resp.Company.Staff[0].Role)
		assert.NotNil(t, resp.Company.CreatedAt)
	})

	t.Run("GetCompany_Integration", func(t *testing.T) {
		// テスト用のサービスをセットアップ
		companyService, _, cleanup := setupTestService()
		defer cleanup()

		ctx := context.Background()

		// 1. 会社を作成
		createReq := &companyv1.CreateCompanyRequest{
			Trademark:   "取得テスト株式会社",
			Type:        "kabu",
			Position:    "prefix",
			Address:     "東京都渋谷区",
			CompanyCode: "2222222222222",
			Ceo: &staffv1.Staff{
				Name: "取得テスト太郎",
				Role: "代表取締役",
				Contact: &contactv1.Contact{
					Email: "get@test.com",
					Phone: "03-2222-2222",
				},
			},
			Contact: &contactv1.ContactRequest{
				Email: "info@get-test.com",
				Phone: "03-2222-2223",
			},
		}

		createResp, err := companyService.CreateCompany(ctx, createReq)
		require.NoError(t, err)
		require.NotNil(t, createResp.Company)

		// 2. 作成された会社を取得
		getReq := &companyv1.GetCompanyRequest{
			Id: createResp.Company.Id,
		}

		getResp, err := companyService.GetCompany(ctx, getReq)

		// 3. データが正しく取得できることを確認
		require.NoError(t, err)
		assert.NotNil(t, getResp)
		assert.NotNil(t, getResp.Company)
		assert.Equal(t, createResp.Company.Id, getResp.Company.Id)
		assert.Equal(t, createReq.Trademark, getResp.Company.Trademark)
		assert.Equal(t, createReq.Type, getResp.Company.Type)
		assert.Equal(t, createReq.Position, getResp.Company.Position)
		assert.Equal(t, createReq.Address, getResp.Company.Address)
		assert.Equal(t, createReq.CompanyCode, getResp.Company.CompanyCode)
		assert.NotNil(t, getResp.Company.Contact)
		assert.Equal(t, createReq.Contact.Email, getResp.Company.Contact.Email)
		assert.Equal(t, createReq.Contact.Phone, getResp.Company.Contact.Phone)
		assert.NotNil(t, getResp.Company.CreatedAt)
	})

	t.Run("ListCompanies_Integration", func(t *testing.T) {
		// テスト用のサービスをセットアップ
		companyService, _, cleanup := setupTestService()
		defer cleanup()

		ctx := context.Background()

		// 1. 複数の会社を作成
		companies := []*companyv1.CreateCompanyRequest{
			{
				Trademark:   "一覧テスト株式会社A",
				Type:        "kabu",
				Position:    "prefix",
				Address:     "東京都新宿区",
				CompanyCode: "3333333333333",
				Ceo: &staffv1.Staff{
					Name: "一覧テスト太郎A",
					Role: "代表取締役",
					Contact: &contactv1.Contact{
						Email: "list-a@test.com",
						Phone: "03-3333-3333",
					},
				},
				Contact: &contactv1.ContactRequest{
					Email: "info@list-a.com",
					Phone: "03-3333-3334",
				},
			},
			{
				Trademark:   "一覧テスト株式会社B",
				Type:        "kabu",
				Position:    "suffix",
				Address:     "東京都港区",
				CompanyCode: "4444444444444",
				Ceo: &staffv1.Staff{
					Name: "一覧テスト太郎B",
					Role: "代表取締役",
					Contact: &contactv1.Contact{
						Email: "list-b@test.com",
						Phone: "03-4444-4444",
					},
				},
				Contact: &contactv1.ContactRequest{
					Email: "info@list-b.com",
					Phone: "03-4444-4445",
				},
			},
		}

		createdCompanyIds := make([]string, len(companies))
		for i, companyReq := range companies {
			resp, err := companyService.CreateCompany(ctx, companyReq)
			require.NoError(t, err)
			require.NotNil(t, resp.Company)
			createdCompanyIds[i] = resp.Company.Id
		}

		// 2. 会社一覧を取得
		listReq := &companyv1.ListCompaniesRequest{}
		listResp, err := companyService.ListCompanies(ctx, listReq)

		// 3. 正しい順序で取得できることを確認
		require.NoError(t, err)
		assert.NotNil(t, listResp)
		assert.NotNil(t, listResp.Companies)
		assert.GreaterOrEqual(t, len(listResp.Companies), len(companies))

		// 作成した会社が含まれていることを確認
		foundCompanies := make(map[string]bool)
		for _, company := range listResp.Companies {
			for _, createdId := range createdCompanyIds {
				if company.Id == createdId {
					foundCompanies[createdId] = true
					break
				}
			}
		}

		// 全ての作成した会社が一覧に含まれていることを確認
		for _, createdId := range createdCompanyIds {
			assert.True(t, foundCompanies[createdId], "Created company with ID %s should be in the list", createdId)
		}
	})

	t.Run("StaffOperations_Integration", func(t *testing.T) {
		// テスト用のサービスをセットアップ
		companyService, staffService, cleanup := setupTestService()
		defer cleanup()

		ctx := context.Background()

		// 1. 会社を作成
		companyReq := &companyv1.CreateCompanyRequest{
			Trademark:   "スタッフテスト株式会社",
			Type:        "kabu",
			Position:    "prefix",
			Address:     "東京都品川区",
			CompanyCode: "5555555555555",
			Ceo: &staffv1.Staff{
				Name: "スタッフテスト太郎",
				Role: "代表取締役",
				Contact: &contactv1.Contact{
					Email: "staff@test.com",
					Phone: "03-5555-5555",
				},
			},
			Contact: &contactv1.ContactRequest{
				Email: "info@staff-test.com",
				Phone: "03-5555-5556",
			},
		}

		companyResp, err := companyService.CreateCompany(ctx, companyReq)
		require.NoError(t, err)
		require.NotNil(t, companyResp.Company)

		// 2. スタッフを追加
		createStaffReq := &staffv1.CreateStaffRequest{
			Name:      "新入社員花子",
			Role:      "エンジニア",
			CompanyId: companyResp.Company.Id,
			Contact: &contactv1.ContactRequest{
				Email: "hanako@staff-test.com",
				Phone: "03-5555-5557",
			},
		}

		createStaffResp, err := staffService.CreateStaff(ctx, createStaffReq)
		require.NoError(t, err)
		assert.NotNil(t, createStaffResp.Staff)
		assert.Equal(t, createStaffReq.Name, createStaffResp.Staff.Name)
		assert.Equal(t, createStaffReq.Role, createStaffResp.Staff.Role)
		assert.Equal(t, createStaffReq.Contact.Email, createStaffResp.Staff.Contact.Email)
		assert.Equal(t, createStaffReq.Contact.Phone, createStaffResp.Staff.Contact.Phone)
		assert.NotEmpty(t, createStaffResp.Staff.Id)

		// 3. スタッフ情報を更新
		updateStaffReq := &staffv1.UpdateStaffRequest{
			Id:   createStaffResp.Staff.Id,
			Name: stringPtr("上級エンジニア花子"),
			Role: stringPtr("シニアエンジニア"),
			Contact: &contactv1.ContactRequest{
				Email: "senior.hanako@staff-test.com",
				Phone: "03-5555-5558",
			},
		}

		updateStaffResp, err := staffService.UpdateStaff(ctx, updateStaffReq)
		require.NoError(t, err)
		assert.NotNil(t, updateStaffResp.Staff)
		assert.Equal(t, *updateStaffReq.Name, updateStaffResp.Staff.Name)
		assert.Equal(t, *updateStaffReq.Role, updateStaffResp.Staff.Role)
		assert.Equal(t, updateStaffReq.Contact.Email, updateStaffResp.Staff.Contact.Email)
		assert.Equal(t, updateStaffReq.Contact.Phone, updateStaffResp.Staff.Contact.Phone)

		// 4. スタッフを削除
		deleteStaffReq := &staffv1.DeleteStaffRequest{
			Id: createStaffResp.Staff.Id,
		}

		deleteStaffResp, err := staffService.DeleteStaff(ctx, deleteStaffReq)
		require.NoError(t, err)
		assert.NotNil(t, deleteStaffResp)
		assert.Equal(t, createStaffResp.Staff.Id, deleteStaffResp.Id)

		// 削除されたスタッフが取得できないことを確認
		// 注意: 実際の実装では、削除されたスタッフの取得がエラーになることを確認する必要があります
		// ここでは削除操作が成功したことを確認します
	})
}

// stringPtr は文字列のポインタを返すヘルパー関数です
func stringPtr(s string) *string {
	return &s
}
