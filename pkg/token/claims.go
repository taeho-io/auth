package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/taeho-io/auth"
)

type Claims struct {
	jwt.StandardClaims

	TokenType auth.TokenType `json:"token_type"`
	UserID    int64          `json:"user_id"`
}
