package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
)

func TestParseHandler(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	accessToken, _ := tokenSvc.NewAccessToken(token.Claims{UserID: testUserId})
	req := &auth.ParseRequest{
		AccessToken: accessToken,
	}
	res, err := Parse(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestParseHandler_Error(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	req := &auth.ParseRequest{
		AccessToken: "invalid_token",
	}
	res, err := Parse(tokenSvc)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}
