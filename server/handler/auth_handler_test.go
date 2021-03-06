package handler

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/auth/pkg/token"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	testUserId                      = int64(1234)
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

func TestAuthHandler_Validate_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		UserId: 0,
	}
	tkn := token.Mock()

	_, err := Auth(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
	assert.Error(t, err, status.Error(codes.InvalidArgument, err.Error()))
}

func TestAuthHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		UserId: testUserId,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := token.NewMockToken(ctrl)
	tkn.
		EXPECT().
		NewAccessToken(token.Claims{UserID: testUserId}).
		Return("", errors.New("failed"))

	_, err := Auth(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
}

func TestAuthHandler_NewRefreshToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		UserId: testUserId,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := token.NewMockToken(ctrl)
	tkn.
		EXPECT().
		NewAccessToken(token.Claims{UserID: testUserId}).
		Return("token", nil)
	tkn.
		EXPECT().
		NewRefreshToken(token.Claims{UserID: testUserId}).
		Return("", errors.New("failed"))

	_, err := Auth(testAccessTokenExpiringDuration, tkn)(ctx, req)
	assert.NotNil(t, err)
}
