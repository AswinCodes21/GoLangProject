package repository

import (
	"errors"
	"my_project/internal/entity"

	"github.com/jmoiron/sqlx"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CreateUser(user *entity.User) (*entity.User, error) {
	query := `INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, user.FullName, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `SELECT id, full_name, email, password FROM users WHERE email = $1`
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
