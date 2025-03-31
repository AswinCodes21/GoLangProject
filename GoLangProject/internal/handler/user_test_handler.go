package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"my_project/internal/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Signup(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUseCase) Login(email, password string) (*entity.User, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUseCase) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestUserHandlers(t *testing.T) {
	mockUseCase := new(MockUserUseCase)

	t.Run("Signup Success", func(t *testing.T) {
		handler := NewSignupHandler(mockUseCase)

		user := &entity.User{
			FullName: "John Doe",
			Email:    "john@example.com",
			Password: "hashedpassword",
		}

		mockUseCase.On("Signup", mock.Anything).Return(user, nil)

		requestBody, _ := json.Marshal(entity.User{
			FullName: "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		})

		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.Signup(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Signup Failure - Email already exists", func(t *testing.T) {
		handler := NewSignupHandler(mockUseCase)

		mockUseCase.On("Signup", mock.Anything).Return(nil, errors.New("email already exists"))

		requestBody, _ := json.Marshal(entity.User{
			FullName: "John Doe",
			Email:    "existing@example.com",
			Password: "password123",
		})

		req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.Signup(rec, req)

		assert.Equal(t, http.StatusConflict, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Login Success", func(t *testing.T) {
		handler := NewLoginHandler(mockUseCase)

		user := &entity.User{
			FullName: "John Doe",
			Email:    "john@example.com",
			Password: "hashedpassword",
		}

		mockUseCase.On("Login", "john@example.com", "password123").Return(user, nil)

		requestBody, _ := json.Marshal(LoginRequest{
			Email:    "john@example.com",
			Password: "password123",
		})

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.Login(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("Login Failure - Invalid credentials", func(t *testing.T) {
		handler := NewLoginHandler(mockUseCase)

		mockUseCase.On("Login", "wrong@example.com", "wrongpassword").Return(nil, errors.New("invalid email or password"))

		requestBody, _ := json.Marshal(LoginRequest{
			Email:    "wrong@example.com",
			Password: "wrongpassword",
		})

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		handler.Login(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		mockUseCase.AssertExpectations(t)
	})
}
