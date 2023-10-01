package dto

import "time"

// エンティティやリポジトリのメソッドを組み合わせてビジネスロジックを作り上げるのがサービスの責務
// サービスはエンティティを扱うが、上位レイヤーに対してはDTO（Data Transfer Object）に変換してから渡します。
// これは、アプリケーション層にドメイン知識が流出するのを防ぐため。
// DTOを使わなかったらドメインロジックがアプリケーションとサービスに散らばってしまい、影響範囲が大きくなってしまう。
// DTOはアプリケーション層に特化しておりバリデーションやリクエストレスポンスとのバインドにも使用される。
// さらに、エンティティからの変換の際にアプリケーション層に後悔したくないフィールドを制限したり整形したりといったことも可能。

// IDのみユーザーデータ
type UserIDModel struct {
	ID uint `query:"id" form:"id" validate:"required"`
}

// ユーザーデータ
type UserModel struct {
	ID        uint      `json:"id" form:"id" validate:"required"`
	Name      string    `json:"name" form:"name" validate:"required,max=32"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewUserModel(id uint, name string, createdAt time.Time, updatedAt time.Time) *UserModel {
	return &UserModel{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
