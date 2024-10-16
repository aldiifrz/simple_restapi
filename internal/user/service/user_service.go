package service

import (
	"simple_restapi/internal/user/entity"
	"simple_restapi/internal/user/repository"
	"time"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *entity.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserById(id int64) (*entity.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) UpdateUser(user *entity.User) error {
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
