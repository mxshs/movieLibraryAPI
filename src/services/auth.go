package services

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	signingKey []byte
}

func NewAuthService() *AuthService {
	return &AuthService{[]byte("Test")}
}

func (as *AuthService) CreateTokenPair(role domain.Role) (*domain.UserTokenPair, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.TokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(4 * time.Hour).Unix(),
		},
		Role: role,
	}).SignedString(as.signingKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.TokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
		},
		Role: role,
	}).SignedString(as.signingKey)
	if err != nil {
		return nil, err
	}

	return &domain.UserTokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AuthService) UseRefreshToken(refreshToken string, claim *domain.TokenClaim) (*domain.UserTokenPair, error) {
	if err := as.ValidateToken(refreshToken, claim); err != nil {
		return nil, err
	}

	return as.CreateTokenPair(claim.Role)
}

func (as *AuthService) ValidateToken(token string, claim *domain.TokenClaim) error {
	tok, err := jwt.ParseWithClaims(
		token,
		&domain.TokenClaim{},
		func(tok *jwt.Token) (any, error) {
			return []byte("Test"), nil
		},
	)
	if err != nil {
		return err
	}

	if claims, ok := tok.Claims.(*domain.TokenClaim); ok {
		if claims.Role >= claim.Role {
			return nil
		}
		return fmt.Errorf("invalid token")
	}

	return fmt.Errorf("invalid token")
}
