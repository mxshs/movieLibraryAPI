package adaptermodels

import "mxshs/movieLibrary/src/domain"

// User example
type User struct {
	Username string      `json:"username" example:"test_user" validate:"required"`
	Password string      `json:"password" example:"test_password" validate:"required"`
	Role     domain.Role `json:"role" example:"0" format:"domain.Role"`
}

// UserUpdate example
type UserUpdate struct {
	Username    string `json:"username" example:"test_user" validate:"required"`
	Password    string `json:"password" example:"test_password" validate:"required"`
	NewUsername string `json:"new_username" example:"test_user_new"`
	NewPassword string `json:"new_password" example:"test_password_new"`
}

// UserLogin example
type UserLogin struct {
	Username string `json:"username" example:"test_user" validate:"required"`
	Password string `json:"password" example:"test_password" validate:"required"`
}
