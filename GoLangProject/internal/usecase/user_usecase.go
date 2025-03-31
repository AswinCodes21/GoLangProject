package usecase

import (
	"errors"
	"my_project/internal/entity"
	"my_project/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) Signup(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	return u.userRepo.CreateUser(user)
}

func (u *UserUseCase) Login(email, password string) (*entity.User, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (u *UserUseCase) GetUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.GetUserByEmail(email)
}
