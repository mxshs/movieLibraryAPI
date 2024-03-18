package api

import (
	"encoding/json"
	"mxshs/movieLibrary/src/domain"
	adaptermodels "mxshs/movieLibrary/src/handlers/adapter_models"
	"mxshs/movieLibrary/src/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
	authService *services.AuthService
}

func NewUserHandler(userService *services.UserService, authService *services.AuthService) *UserHandler {
	return &UserHandler{userService, authService}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user adaptermodels.User

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = uh.userService.CreateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokens, err := uh.authService.CreateTokenPair(domain.ADM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user adaptermodels.UserResetPassword

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = uh.userService.LoginUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = uh.userService.UpdateUser(user.Username, user.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user adaptermodels.User

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	uh.userService.DeleteUser(user.Username)
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user adaptermodels.UserLogin

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := uh.userService.LoginUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokens, err := uh.authService.CreateTokenPair(u.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(tokens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}
