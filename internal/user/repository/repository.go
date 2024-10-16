package repository

import "simple_restapi/internal/user/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserById(id int64) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id int64) error
}
