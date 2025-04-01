package handler

import (
	"encoding/json"
	"my_project/internal/usecase"
	"net/http"
)

type GetUsersHandler struct {
	userUseCase usecase.UserUseCaseInterface
}

func NewGetUsersHandler(userUseCase usecase.UserUseCaseInterface) *GetUsersHandler {
	return &GetUsersHandler{userUseCase: userUseCase}
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *GetUsersHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		http.Error(w, "Error retrieving users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
			Password: user.Password,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userResponses); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}
