package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/manlikeNacho/Sissors/src/repository/sliceRepo"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid token signing")
	ErrInvalidToken         = errors.New("invalid token provided")
	ErrParsingToken         = errors.New("error parsing token")
)

type TokenServiceInterface interface {
	NewTokenService(tokenSecret string)
	Signout(userid string) error
	ValidateAccessToken(tokenString string) (*accessTokenCustomClaims, error)
}

type tokenService struct {
	TokenSecret       string
	AccessExpiration  time.Duration
	RefreshExpiration time.Duration
}

type accessTokenCustomClaims struct {
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

// tokenservice implements TokenServiceInterface
var TokenServ TokenServiceInterface = &tokenService{}

func (t *tokenService) NewTokenService(tokenSecret string) {
	t.TokenSecret = tokenSecret
}

func (t *tokenService) Signout(userid string) error {
	if err := sliceRepo.TokenRepo.DeleteUserRefreshToken(userid); err != nil {
		return err
	}
	return nil
}

func (t *tokenService) ValidateAccessToken(tokenString string) (*accessTokenCustomClaims, error) {
	claims := &accessTokenCustomClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(t.TokenSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, ErrParsingToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	accessTokenClaims, ok := token.Claims.(*accessTokenCustomClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return accessTokenClaims, nil
}
