package handler

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
)

func TestJWKSHandler(t *testing.T) {
	ctx := context.Background()
	verifyingKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(token.MockVerifyPEM))
	assert.Nil(t, err)

	resp, err := Jwks(verifyingKey)(ctx, &auth.JwksRequest{})
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	assert.Len(t, resp.Keys, 1)
	assert.Equal(t, resp.Keys[0].Kty, "RSA")
	assert.NotEmpty(t, resp.Keys[0].E)
	assert.NotEmpty(t, resp.Keys[0].N)
}
