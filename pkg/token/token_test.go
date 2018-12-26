package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJWTSigningMethod_RS256(t *testing.T) {
	jwtSigningMethod, err := jwtSigningMethod("RS256")
	assert.Equal(t, jwtSigningMethod, jwt.SigningMethodRS256)
	assert.Nil(t, err)
}

func TestJWTSigningMethod_RS512(t *testing.T) {
	jwtSigningMethod, err := jwtSigningMethod("RS512")
	assert.Equal(t, jwtSigningMethod, jwt.SigningMethodRS512)
	assert.Nil(t, err)
}

func TestInvalidJWTSigningMethod(t *testing.T) {
	jwtSigningMethod, err := jwtSigningMethod("invalid")
	assert.Nil(t, jwtSigningMethod)
	assert.NotNil(t, err)
}

func TestNewAccessToken_InvalidSigningMethod(t *testing.T) {
	cfg := NewConfig(
		"invalid",
		MockSigningPEM,
		MockVerifyPEM,
		"MOCK_TOKEN_ISSUER",
		time.Hour,
		time.Hour*24*365,
	)
	claims := Claims{
		UserID: 1234,
	}
	tokenSvc, err := New(cfg)
	assert.Nil(t, err)
	token, err := tokenSvc.NewAccessToken(claims)
	assert.NotNil(t, err)
	assert.Empty(t, token, "")
}

func TestNewAccessToken(t *testing.T) {
	claims := Claims{
		UserID: 1234,
	}
	tokenSvc := Mock()
	token, err := tokenSvc.NewAccessToken(claims)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestNewRefreshToken_InvalidSigningMethod(t *testing.T) {
	cfg := NewConfig(
		"invalid",
		MockSigningPEM,
		MockVerifyPEM,
		"MOCK_TOKEN_ISSUER",
		time.Hour,
		time.Hour*24*365,
	)
	claims := Claims{
		UserID: 1234,
	}
	tokenSvc, err := New(cfg)
	assert.Nil(t, err)
	token, err := tokenSvc.NewRefreshToken(claims)
	assert.NotNil(t, err)
	assert.Empty(t, token, "")
}

func TestNewRefreshToken(t *testing.T) {
	claims := Claims{
		UserID: 1234,
	}
	tokenSvc := Mock()
	token, err := tokenSvc.NewRefreshToken(claims)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestVerifyToken(t *testing.T) {
	claims := Claims{
		UserID: 1234,
	}
	tokenSvc := Mock()
	token, _ := tokenSvc.NewAccessToken(claims)
	err := tokenSvc.VerifyToken(token)
	assert.Nil(t, err)
}

func TestParseToken_SignedWithInvalidSigningMethod(t *testing.T) {
	hs256Token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDU3NjMxNTQsImp0aSI6ImJnaDZtZ25r" +
		"a2k2OWtqOGY2OHFnIiwiaWF0IjoxNTQ1NzU5NTU0LCJpc3MiOiJhdXRoLnRhZWhvLmlvIiwidXNlcl9pZCI6MTIzNH0" +
		".JnqyHGQ2ax1dQJxjAX8qagkG8wZZf5V0U4UN8MM_PhVHHcyyfH9e1UJDJfhxj2VZT1MklvgXbR_7I6B2q4SIKg"
	tokenSvc := Mock()
	_, err := tokenSvc.ParseToken(hs256Token)
	assert.NotNil(t, err)
}

func TestParseToken(t *testing.T) {
	claims := Claims{
		UserID: 1234,
	}
	tokenSvc := Mock()
	token, _ := tokenSvc.NewAccessToken(claims)
	claims, err := tokenSvc.ParseToken(token)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
}

func TestParseToken_Error(t *testing.T) {
	tokenSvc := Mock()
	token := "invalid_token"
	claims, err := tokenSvc.ParseToken(token)
	assert.NotNil(t, err)
	assert.Empty(t, claims)
}
