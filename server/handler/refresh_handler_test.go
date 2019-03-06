package handler

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth/pkg/token"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
)

func TestRefreshHandler(t *testing.T) {
	ctx := context.Background()

	tkn := token.Mock()

	refreshToken, _ := tkn.NewRefreshToken(token.Claims{UserID: testUserId})
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	res, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRefreshHandler_Validate_Error(t *testing.T) {
	ctx := context.Background()

	tkn := token.Mock()

	req := &auth.RefreshRequest{
		RefreshToken: "invalid_token",
	}
	res, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestRefreshHandler_InvalidToken_Error(t *testing.T) {
	ctx := context.Background()

	tkn := token.Mock()

	req := &auth.RefreshRequest{
		RefreshToken: "invalid_token_with_dummy_to_make_its_length_bigger_than_30",
	}
	res, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestRefreshHandler_InvalidTokenType_Error(t *testing.T) {
	ctx := context.Background()

	refreshToken, _ := token.Mock().NewRefreshToken(token.Claims{UserID: testUserId})
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := token.NewMockToken(ctrl)
	tkn.
		EXPECT().
		ParseToken(refreshToken).
		Return(token.Claims{UserID: testUserId, TokenType: auth.TokenType_ACCESS_TOKEN}, nil)

	res, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestRefreshHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()

	refreshToken, _ := token.Mock().NewRefreshToken(token.Claims{UserID: testUserId})
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := token.NewMockToken(ctrl)
	tkn.
		EXPECT().
		ParseToken(refreshToken).
		Return(token.Claims{UserID: testUserId, TokenType: auth.TokenType_REFRESH_TOKEN}, nil)
	tkn.
		EXPECT().
		NewAccessToken(token.Claims{UserID: testUserId}).
		Return("", errors.New("failed"))

	res, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
}
