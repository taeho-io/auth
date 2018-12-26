package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
)

func TestVerifyHandler(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	accessToken, _ := tokenSvc.NewAccessToken(token.Claims{UserID: testUserId})
	req := &auth.VerifyRequest{
		AccessToken: accessToken,
	}
	res, err := Verify(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.IsValid, true)
}

func TestVerifyHandler_Invalid(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	req := &auth.VerifyRequest{
		AccessToken: "invalid",
	}
	res, err := Verify(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.IsValid, false)
}
