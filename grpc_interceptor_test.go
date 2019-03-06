package auth

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/go-taeho/id"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func TestAuthFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenString := "test_jwt_token_string"

	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("Authorization", "Bearer "+tokenString))

	userID := id.New().Must()
	authCli := NewMockAuthClient(ctrl)
	authCli.
		EXPECT().
		Parse(ctx, &auth.ParseRequest{
			AccessToken: tokenString,
		}).
		Return(&auth.ParseResponse{
			TokenType: auth.TokenType_ACCESS_TOKEN,
			UserId:    userID,
		}, nil)

	ctx, err := authFunc(authCli)(ctx)

	assert.NotNil(t, ctx)
	assert.Nil(t, err)

	assert.Equal(t, metautils.ExtractIncoming(ctx).Get(xTokenType), fmt.Sprintf("%d", auth.TokenType_ACCESS_TOKEN))
	assert.Equal(t, metautils.ExtractIncoming(ctx).Get(xTokenUserID), fmt.Sprintf("%v", userID))
	assert.Equal(t, metautils.ExtractOutgoing(ctx).Get(xTokenType), fmt.Sprintf("%d", auth.TokenType_ACCESS_TOKEN))
	assert.Equal(t, metautils.ExtractOutgoing(ctx).Get(xTokenUserID), fmt.Sprintf("%v", userID))
}

func TestUserIDFromIncomingContext(t *testing.T) {
	userID := id.New().Must()
	m := make(map[string]string)
	m[xTokenUserID] = fmt.Sprintf("%v", userID)
	md := metadata.New(m)

	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, md)

	userIDFromContext, err := UserIDFromIncomingContext(ctx)
	assert.Equal(t, int64(userID), userIDFromContext)
	assert.Nil(t, err)
}

func TestUserIDFromIncomingContext__through_authFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenString := "test_jwt_token_string"

	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("Authorization", "Bearer "+tokenString))

	userID := id.New().Must()
	authCli := NewMockAuthClient(ctrl)
	authCli.
		EXPECT().
		Parse(ctx, &auth.ParseRequest{
			AccessToken: tokenString,
		}).
		Return(&auth.ParseResponse{
			TokenType: auth.TokenType_ACCESS_TOKEN,
			UserId:    userID,
		}, nil)

	ctx, err := authFunc(authCli)(ctx)
	assert.NotNil(t, ctx)
	assert.Nil(t, err)

	userIDFromContext, err := UserIDFromIncomingContext(ctx)
	assert.Equal(t, userID, userIDFromContext)
	assert.Nil(t, err)
}

func TestVerifyUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenString := "test_jwt_token_string"

	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("Authorization", "Bearer "+tokenString))

	userID := id.New().Must()
	authCli := NewMockAuthClient(ctrl)
	authCli.
		EXPECT().
		Parse(ctx, &auth.ParseRequest{
			AccessToken: tokenString,
		}).
		Return(&auth.ParseResponse{
			TokenType: auth.TokenType_ACCESS_TOKEN,
			UserId:    userID,
		}, nil)

	ctx, err := authFunc(authCli)(ctx)
	assert.NotNil(t, ctx)
	assert.Nil(t, err)

	err = VerifyUser(ctx, userID, false)
	assert.Nil(t, err)

	err = VerifyUser(ctx, userID, true)
	assert.Nil(t, err)

	err = VerifyUser(ctx, userID+1, false)
	assert.Error(t, err)

	err = VerifyUser(ctx, userID+1, true)
	assert.Error(t, err)

	err = VerifyUser(context.Background(), userID, false)
	assert.Error(t, err)

	err = VerifyUser(context.Background(), userID, true)
	assert.Nil(t, err)

	err = VerifyUser(context.Background(), userID+1, false)
	assert.Error(t, err)

	err = VerifyUser(context.Background(), userID+1, true)
	assert.Nil(t, err)
}
