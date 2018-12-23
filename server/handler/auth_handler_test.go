package handler

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/mocks"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
)

const (
	testUserId                      = "test_user_id"
	testAccessTokenExpiringDuration = time.Hour
)

func TestAuthHandler(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		UserId: testUserId,
	}

	tkn := token.Mock()

	res, err := Auth(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.UserId, testUserId)
	assert.NotNil(t, res.AccessToken)
	assert.NotNil(t, res.RefreshToken)
	assert.Equal(t, res.ExpiresIn, int64(testAccessTokenExpiringDuration.Seconds()))
}

func TestAuthHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		UserId: testUserId,
	}

	tkn := new(mocks.Token)
	tkn.On("NewAccessToken", token.Claims{UserID: testUserId}).
		Return("", errors.New("failed"))
	_, err := Auth(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
}

func TestAuthHandler_NewRefreshToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		UserId: testUserId,
	}

	tkn := new(mocks.Token)
	tkn.On("NewAccessToken", token.Claims{UserID: testUserId}).
		Return("token", nil)
	tkn.On("NewRefreshToken", token.Claims{UserID: testUserId}).
		Return("", errors.New("failed"))
	_, err := Auth(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
}
