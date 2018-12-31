package auth

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/taeho-go/id"
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
		Parse(ctx, &ParseRequest{
			AccessToken: tokenString,
		}).
		Return(&ParseResponse{
			TokenType: TokenType_ACCESS_TOKEN,
			UserId:    userID,
		}, nil)

	ctx, err := authFunc(authCli)(ctx)

	assert.NotNil(t, ctx)
	assert.Nil(t, err)

	assert.Equal(t, metautils.ExtractIncoming(ctx).Get(xTokenType), fmt.Sprintf("%d", TokenType_ACCESS_TOKEN))
	assert.Equal(t, metautils.ExtractIncoming(ctx).Get(xTokenUserID), fmt.Sprintf("%v", userID))
	assert.Equal(t, metautils.ExtractOutgoing(ctx).Get(xTokenType), fmt.Sprintf("%d", TokenType_ACCESS_TOKEN))
	assert.Equal(t, metautils.ExtractOutgoing(ctx).Get(xTokenUserID), fmt.Sprintf("%v", userID))
}
