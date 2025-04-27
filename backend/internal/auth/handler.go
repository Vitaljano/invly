package auth

import (
	"fmt"
	"net/http"

	"github.com/Vitaljano/invly/backend/config"
	"github.com/Vitaljano/invly/backend/pkg/jwt"
	"github.com/Vitaljano/invly/backend/pkg/req"
	"github.com/Vitaljano/invly/backend/pkg/res"
)

type AuthHandlerDeps struct {
	*AuthService
	*config.Config
}

type AuthHandler struct {
	*AuthService
	*config.Config
}

func NewAuthHandler(mux *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
		Config:      deps.Config,
	}

	mux.HandleFunc("POST /auth/registration", handler.Register())
	mux.HandleFunc("POST /auth/login", handler.Login())

}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Println(h.Config.Auth.Secret)
		email, err := h.AuthService.Register(d.Email, d.Name, d.Password)

		token, err := jwt.NewJwt(h.Config.Auth.Secret).Create(jwt.JWTData{Email: email})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponse{
			Token: token,
		}

		res.Json(w, data, http.StatusAccepted)

	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := req.HandleBody[LoginRequest](w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email, err := h.AuthService.Login(d.Email, d.Password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJwt(h.Config.Auth.Secret).Create(jwt.JWTData{Email: email})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		res.Json(w, LoginResponse{
			Token: token,
		}, http.StatusOK)
	}
}
