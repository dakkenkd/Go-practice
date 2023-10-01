package service

import (
	"layered_architecture/domain/dto"
	"layered_architecture/domain/entity"
	"layered_architecture/domain/repository"
)

type IUserService interface {
	FindByID(id uint) (*model.User, error)
	Create(user *dto.UserModel) error
	Update(user *dto.UserModel) error
	Update(user *dto.UserModel) error
	Delete(id uint) error
}

type UserService struct {
	repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo}
}

func (s *UserService) FindByID(id uint) (*dto.UserModel, error) {
	user err := s.IUserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.convertTo(user), nil
}

func (s *UserService) Create(user *dto.UserModel) error {
	u := s.convertFrom(user)
	return s.IUserRepository.Create(u)
}

func (s *UserService) Update(user *dto.UserModel) error {
	u := s.convertFrom(user)
	return s.IUserRepository.Update(u)
}

func (s *UserService) Delete(id uint) error {
	return s.IUserRepository.Delete(id)
}

// エンティティからDTOに変換
func (s *UserService) convertTo(user *entity.User) *dto.UserModel {
	return dto.NewUserModel(user.ID, user.Name, user.CreatedAt, user.UpdatedAt)
}

func (s *UserService) convertFrom(user *dto.UserModel) *entity.User {
	return entity.NewUser(user.ID, user.Name, user.CreatedAt, user.UpdatedAt)
}