package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJWTSigningMethodWithHS256(t *testing.T) {
	jwtSigningMethod, err := jwtSigningMethod("HS256")
	assert.Equal(t, jwtSigningMethod, jwt.SigningMethodHS256)
	assert.Nil(t, err)
}

func TestJWTSigningMethodWithHS512(t *testing.T) {
	jwtSigningMethod, err := jwtSigningMethod("HS512")
	assert.Equal(t, jwtSigningMethod, jwt.SigningMethodHS512)
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
		"MOCK_SIGNING_KEY",
		"MOCK_TOKEN_ISSUER",
		time.Hour,
		time.Hour*24*365,
	)
	claims := Claims{
		UserID: "test_user_id",
	}
	tokenSvc := New(cfg)
	token, err := tokenSvc.NewAccessToken(claims)
	assert.NotNil(t, err)
	assert.Empty(t, token, "")
}

func TestNewAccessToken(t *testing.T) {
	claims := Claims{
		UserID: "test_user_id",
	}
	tokenSvc := Mock()
	token, err := tokenSvc.NewAccessToken(claims)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestNewRefreshToken_InvalidSigningMethod(t *testing.T) {
	cfg := NewConfig(
		"invalid",
		"MOCK_SIGNING_KEY",
		"MOCK_TOKEN_ISSUER",
		time.Hour,
		time.Hour*24*365,
	)
	claims := Claims{
		UserID: "test_user_id",
	}
	tokenSvc := New(cfg)
	token, err := tokenSvc.NewRefreshToken(claims)
	assert.NotNil(t, err)
	assert.Empty(t, token, "")
}

func TestNewRefreshToken(t *testing.T) {
	claims := Claims{
		UserID: "test_user_id",
	}
	tokenSvc := Mock()
	token, err := tokenSvc.NewRefreshToken(claims)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestValidateToken(t *testing.T) {
	claims := Claims{
		UserID: "test_user_id",
	}
	tokenSvc := Mock()
	token, _ := tokenSvc.NewAccessToken(claims)
	err := tokenSvc.ValidateToken(token)
	assert.Nil(t, err)
}

func TestParseToken_SignedWithInvalidSigningMethod(t *testing.T) {
	rs256Token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4" +
		"gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.TCYt5XsITJX1CxPCT8yAV-TVkIEq_PbChOMqsLfRo" +
		"Psnsgw5WEuts01mq-pQy7UJiN5mgRxD-WUcX16dUEMGlv50aqzpqh4Qktb3rk-BuQy72IFLOqV0G_zS245-kronKb7" +
		"8cPN25DGlcTwLtjPAYuNzVBAh4vGHSrQyHUdBBPM"
	tokenSvc := Mock()
	_, err := tokenSvc.ParseToken(rs256Token)
	assert.NotNil(t, err)
}

func TestParseToken(t *testing.T) {
	claims := Claims{
		UserID: "test_user_id",
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
