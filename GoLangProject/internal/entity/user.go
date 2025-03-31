package entity

import "github.com/go-playground/validator/v10"

type User struct {
	ID       int    `db:"id"`
	FullName string `db:"full_name" validate:"required,min=3,max=100"`
	Email    string `db:"email" validate:"required,email"`
	Password string `db:"password" validate:"required,min=6"`
}

func (u *User) ValidateUser() error {
	validate := validator.New()
	return validate.Struct(u)
}
