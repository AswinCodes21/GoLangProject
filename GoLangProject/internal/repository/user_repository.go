package repository

import "my_project/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
