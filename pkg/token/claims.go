package token

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims

	UserID string `json:"user_id"`
}
