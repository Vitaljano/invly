package auth

import (
	"fmt"
	"net/http"

	"github.com/Vitaljano/invly/backend/pkg/res"
)

type AuthHandler struct {
}

func NewAuthHandler(mux *http.ServeMux) {
	handler := &AuthHandler{}

	mux.HandleFunc("/auth/registration", handler.Register())
	mux.HandleFunc("/auth/login", handler.Login())

}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register")
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, LoginResponse{
			Token: "asdasdasdasd",
		}, http.StatusOK)
	}
}
