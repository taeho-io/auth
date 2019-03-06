package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth/pkg/token"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
)

func TestParseHandler(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	accessToken, _ := tokenSvc.NewAccessToken(token.Claims{UserID: 1234})
	req := &auth.ParseRequest{
		AccessToken: accessToken,
	}
	res, err := Parse(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestParseHandler_Validate_Error(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	req := &auth.ParseRequest{
		AccessToken: "invalid_token",
	}
	res, err := Parse(tokenSvc)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestParseHandler_InvalidToken_Error(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	req := &auth.ParseRequest{
		AccessToken: "invalid_token_with_dummy_to_make_its_length_bigger_than_30",
	}
	res, err := Parse(tokenSvc)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}
