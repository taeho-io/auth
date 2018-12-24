package handler

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/mocks"
	"github.com/taeho-io/auth/pkg/token"
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

func TestRefreshHandler_Error_InvalidRefreshToken(t *testing.T) {
	ctx := context.Background()

	tkn := token.Mock()

	req := &auth.RefreshRequest{
		RefreshToken: "invalid_token",
	}
	res, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestRefreshHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()

	refreshToken, _ := token.Mock().NewRefreshToken(token.Claims{UserID: testUserId})
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := mocks.NewMockToken(ctrl)
	tkn.
		EXPECT().
		ParseToken(refreshToken).
		Return(token.Claims{UserID: testUserId}, nil)
	tkn.
		EXPECT().
		NewAccessToken(token.Claims{UserID: testUserId}).
		Return("", errors.New("failed"))

	_, err := Refresh(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
}
