package api

import (
	"encoding/json"
	"fmt"
	"mxshs/movieLibrary/src/domain"
	adaptermodels "mxshs/movieLibrary/src/handlers/adapter_models"
	"mxshs/movieLibrary/src/services"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
	authService *services.AuthService
}

func NewUserHandler(userService *services.UserService, authService *services.AuthService) *UserHandler {
	return &UserHandler{userService, authService}
}

// CreateUser
//
//	@Summary		Create a new user
//	@Description	Creates a new user and returns access+refresh tokens for future auth
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		adaptermodels.User	true	"New user"
//	@Success		200		{object}	domain.UserTokenPair
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/users/ [post]
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

// GetUsers
//
//	@Summary	Get a user by id
//	@Tags		users
//	@Produce	json
//	@Security	Bearer
//	@Param		id	path		int	true	"user id"	minimum(0)
//	@Success	200	{object}	domain.User
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/users/{id} [get]
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	user, err := uh.userService.GetUser(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

// GetUsers
//
//	@Summary	Get all users
//	@Tags		users
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{array}	domain.User
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/users [get]
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

// UpdateUser
//
//	@Summary		Update a user
//	@Description	Update user's username/password, other changes are currently not supported
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"user id"	minimum(0)
//	@Param			user	body		adaptermodels.UserUpdate	true	"Login credentials and fields to update"
//	@Success		200		{object}	domain.UserTokenPair
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/users/{id}/ [patch]
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	var user adaptermodels.UserUpdate

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = uh.userService.LoginUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = uh.userService.UpdateUser(uid, user.NewUsername, user.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// DeleteUser
//
//	@Summary	Delete a user
//	@Tags		users
//	@Security	Bearer
//	@Param		id	path	int	true	"user id"	minimum(0)
//	@Success	200
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/users/{id}/ [delete]
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	uh.userService.DeleteUser(uid)
}

// LoginUser
//
//	@Summary		Login a user
//	@Description	Login a user with username and password and return new token pair
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		domain.User	true	"Login credentials"
//	@Success		200		{object}	domain.UserTokenPair
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/users/login/ [post]
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
