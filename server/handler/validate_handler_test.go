package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
)

func TestValidateHandler(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	accessToken, _ := tokenSvc.NewAccessToken(token.Claims{UserID: testUserId})
	req := &auth.ValidateRequest{
		AccessToken: accessToken,
	}
	res, err := Validate(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.IsValid, true)
}

func TestValidateHandler_Invalid(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	req := &auth.ValidateRequest{
		AccessToken: "invalid",
	}
	res, err := Validate(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.IsValid, false)
}
