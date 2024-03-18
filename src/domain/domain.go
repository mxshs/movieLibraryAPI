package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Actor struct {
	Id   int
	Name string
	// could've used an enum for gender, but its not obvious how to keep values ubiquitous
	Gender    string
	Birthdate time.Time
	Movies    []*Movie
}

type Movie struct {
	Id          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      uint8
	Stars       []*Actor
}

type User struct {
	Id       int
	Username string
	Password string
	Role     Role
}

type Role int

// higher the role value, higher the status
const (
	USR Role = iota
	ADM
)

type TokenClaim struct {
	jwt.StandardClaims
	Role Role
}

type UserTokenPair struct {
	AccessToken  string
	RefreshToken string
}
