package entity

import "time"

type User struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id uint, name string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
