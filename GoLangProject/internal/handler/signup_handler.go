package handler

import (
	"encoding/json"
	"my_project/internal/entity"
	"my_project/internal/usecase"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignupHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewSignupHandler(userUseCase *usecase.UserUseCase) *SignupHandler {
	return &SignupHandler{userUseCase: userUseCase}
}

type SignupRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *SignupHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := &entity.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.ValidateUser(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	createdUser, err := h.userUseCase.Signup(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
