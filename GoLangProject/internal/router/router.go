package router

import (
	"my_project/internal/handler"
	"my_project/internal/repository"
	"my_project/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	signupHandler := handler.NewSignupHandler(userUseCase)
	loginHandler := handler.NewLoginHandler(userUseCase)
	getUsersHandler := handler.NewGetUsersHandler(userUseCase)

	r.HandleFunc("/signup", signupHandler.Signup).Methods("POST")
	r.HandleFunc("/login", loginHandler.Login).Methods("POST")
	r.HandleFunc("/users", getUsersHandler.GetUsers).Methods("GET")
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	return r
}
