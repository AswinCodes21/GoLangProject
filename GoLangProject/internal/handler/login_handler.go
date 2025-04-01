package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"my_project/internal/usecase"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	userUseCase usecase.UserUseCaseInterface
}

func NewLoginHandler(userUseCase usecase.UserUseCaseInterface) *LoginHandler {
	return &LoginHandler{userUseCase: userUseCase}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{"Invalid request body"})
		return
	}

	user, err := h.userUseCase.GetUserByEmail(req.Email)
	if err != nil {
		writeJSONResponse(w, http.StatusUnauthorized, ErrorResponse{"User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("Password comparison failed for user %s: %v", req.Email, err)
		writeJSONResponse(w, http.StatusUnauthorized, ErrorResponse{"Invalid credentials"})
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{"Failed to generate token"})
		return
	}

	writeJSONResponse(w, http.StatusOK, LoginResponse{Token: token})
}

func generateToken(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("Warning: JWT_SECRET is not set, using default secret")
		secret = "default_secret"
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
