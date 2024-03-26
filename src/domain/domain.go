package domain

import (
	"mxshs/movieLibrary/src/utils"

	"github.com/golang-jwt/jwt"
)

// Actor example
type Actor struct {
	Id        int        `json:"id" example:"0" format:"int64"`
	Name      string     `json:"name" example:"Leonardo DiCaprio" validate:"required"`
	Gender    string     `json:"gender" example:"male" validate:"required"`
	Birthdate utils.Date `json:"birthdate" example:"11.11.1974" format:"date"`
}

// ActorDetail example
type ActorDetail struct {
	Id        int        `json:"id" example:"0" format:"int64"`
	Name      string     `json:"name" example:"Leonardo DiCaprio"`
	Gender    string     `json:"gender" example:"male"`
	Birthdate utils.Date `json:"birthdate" example:"11.11.1974" format:"date"`
	Movies    []*Movie   `json:"movies" format:"domain.Movie"`
}

// Movie example
type Movie struct {
	Id          int        `json:"id" example:"0" format:"int64"`
	Title       string     `json:"title" example:"The Wolf of Wall-Street"`
	Description string     `json:"description" example:"Movie about some stuff" validate:"required" maxLength:"1500"`
	ReleaseDate utils.Date `json:"release_date" example:"25.12.2013" format:"date" validate:"required"`
	Rating      uint8      `json:"rating" example:"8" format:"uint8" validate:"required" minimum:"0" maximum:"10"`
}

// MovieDetail example
type MovieDetail struct {
	Id          int        `json:"id" example:"0" format:"int64"`
	Title       string     `json:"title" example:"The Wolf of Wall-Street"`
	Description string     `json:"description" example:"Movie about some stuff" validate:"required" maxLength:"1500"`
	ReleaseDate utils.Date `json:"release_date" example:"25.12.2013" format:"date" validate:"required"`
	Rating      uint8      `json:"rating" example:"8" format:"uint8" validate:"required" minimum:"0" maximum:"10"`
	Stars       []*Actor   `json:"stars" format:"domain.Actor"`
}

// User example
type User struct {
	Id       int    `json:"id" example:"0" format:"int64"`
	Username string `json:"username" example:"test_user"`
	Role     Role   `json:"role" example:"0" format:"domain.Role"`
}

// UserDetail example
type UserDetail struct {
	Id       int    `json:"id" example:"0" format:"int64"`
	Username string `json:"username" example:"test_user"`
	Password string `json:"password" example:"test_password"`
	Role     Role   `json:"role" example:"0" format:"domain.Role"`
}

type Role int

// higher the role value, higher the status
// Role example
const (
	USR Role = iota
	ADM
)

type TokenClaim struct {
	jwt.StandardClaims
	Role Role
}

// UserTokenPair example
type UserTokenPair struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzMTA0NDMsIlJvbGUiOjF9.W7yqNX39GkHzYZnzw6U7gaMib4lmdpipIRXzvSNUfII"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzMjQ4NDMsIlJvbGUiOjF9.2wS0BonGkE-Xa1CaTXSr_OYQYLgtginJYLkk-2n8b_Y"`
}
