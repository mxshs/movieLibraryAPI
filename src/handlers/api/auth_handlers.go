package api

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/services"
	"net/http"
	"strings"
)

type AuthHandler struct {
	as *services.AuthService
}

func NewAuthHandler(as *services.AuthService) *AuthHandler {
	return &AuthHandler{as}
}

func (am *AuthHandler) Authenticate(next http.HandlerFunc, role domain.Role) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if len(tokenHeader) == 0 {
			http.Error(w, "Auth token missing", http.StatusUnauthorized)
			return
		}

		token := strings.Split(tokenHeader, " ")
		if len(token) != 2 {
			http.Error(w, "Wrong auth header value", http.StatusUnauthorized)
			return
		}

		err := am.as.ValidateToken(token[1], &domain.TokenClaim{Role: role})
		if err != nil {
			fmt.Println("i failed")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
