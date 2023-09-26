package repository

import "Go-practice/infrastucture/mysql/entity"

type IUserRepository interface {
	FindById(id uint) (*entity.User, error)
}
