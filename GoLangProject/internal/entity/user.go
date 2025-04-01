package entity

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID       int    `db:"id"`
	FullName string `db:"full_name" validate:"required,min=3,max=100"`
	Email    string `db:"email" validate:"required,email"`
	Password string `db:"password" validate:"required,min=6"`
}

func passwordComplexity(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	uppercase := regexp.MustCompile(`[A-Z]`)
	lowercase := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`[0-9]`)

	return uppercase.MatchString(password) && lowercase.MatchString(password) && digit.MatchString(password)
}

func (u *User) ValidateUser() error {
	validate := validator.New()

	validate.RegisterValidation("password_complexity", passwordComplexity)

	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New("Validation failed: " + err.StructNamespace() + " " + err.Tag())
		}
	}
	return nil
}
