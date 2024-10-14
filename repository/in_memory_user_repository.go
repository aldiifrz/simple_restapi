package repository

import (
	"fmt"
	"simple_restapi/entity"
)

type InMemoryUserRepository struct {
	users map[int64]*entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int64]*entity.User),
	}
}

func (r *InMemoryUserRepository) CreateUser(user *entity.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetUserById(id int64) (*entity.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) UpdateUser(user *entity.User) error {
	_, exists := r.users[user.ID]
	if !exists {
		return fmt.Errorf("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) DeleteUser(id int64) error {
	delete(r.users, id)
	return nil
}
