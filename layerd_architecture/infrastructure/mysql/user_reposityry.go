package mysql

import (
	"database/sql"
	"layered_architecture/domain/repository"
	"layered_architecture/infrastructure/mysql/entity"
)

type UserRepository struct {
	*sql.DB
}

// データベースのコネクションを外側から渡せるようにすることでテストが容易になる

// また、関数の戻り値をインターフェース型にして構造体をreturnすると肩チェックが行われる。
// 構造体がインターフェースを満たしているかどうかをコンパイル時にチェックできるため便利

// IUserRepositoryを満たす構造体を返す(IUserRepositoryはdomainで定義されている)

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	stmt, err := r.Prepare(`SELECT id, name, created_at, updated_at FROM users WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &entity.User{}
	err := stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	stmt, err := r.Prepare(`INSERT INTO users (name, created_at, updated_at) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user *entity.User) error {
	stmt, err := r.Prepare(`UPDATE users SET name = ?, updated_at = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.UpdatedAt, user.ID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id uint) error {
	stmt, err := r.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
