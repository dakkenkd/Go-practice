package mysql

import (
	"layered_architecture/common"
	"layered_architecture/infrastructure/mysql/entity"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_Create(tt *testing.T) {
	tt.Run(
		"正常系：エラーなし",
		func(t *testing.T) {
			user := &entity.User{
				Name: "Alice",
				CreatedAt: common.CurrentTime(),
			}

			// モック用のコネクション作成
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			// SQL、引数、戻り値が意図したものであることを期待する
			mock.ExpectPrepare(`INSERT INTO users`).
				ExpectExec().
				WithArgs(user.Name, user.CreatedAt).
				WillReturnResult(sqlmock.NewResult(1, 1))

			// テスト用のリポジトリを作成
			r := NewUserRepository(db)

			assert.NoError(t, r.CreateUser(user))

			// 上記で指定した通りにモックがよばれることを期待する
			assert.NoError(t, mock.ExpectationsWereMet())
		}
	)
}