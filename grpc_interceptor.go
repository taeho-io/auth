package auth

import (
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	bearerScheme = "bearer"
	xTokenPrefix = "x-token-"
	xTokenUserID = "x-token-user_id"
	xTokenType   = "x-token-type"
)

var (
	ErrInvalidToken = status.Error(codes.Unauthenticated, "invalid token")

	authCli = GetAuthClient()
)

func TokenUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(authFunc(authCli))
}

func authFunc(authCli AuthClient) func(context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		tokenString, err := grpc_auth.AuthFromMD(ctx, bearerScheme)
		if err != nil {
			return ctx, nil
		}

		parseResp, err := authCli.Parse(ctx, &ParseRequest{
			AccessToken: tokenString,
		})
		if err != nil {
			return nil, err
		}
		if parseResp.TokenType != TokenType_ACCESS_TOKEN {
			return nil, ErrInvalidToken
		}

		// Remove all x-token- prefixed metadata since it could be a security threat
		pairs := make([]string, 0)
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key, values := range md {
				if !strings.HasPrefix(strings.ToLower(key), xTokenPrefix) {
					for _, value := range values {
						pairs = append(pairs, key, value)
					}
				}
			}
		}
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(pairs...))

		// Set newly parsed x-token- metadata to both incoming and outgoing
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			xTokenMD := metadata.Pairs(
				xTokenUserID, fmt.Sprintf("%v", parseResp.UserId),
				xTokenType, fmt.Sprintf("%d", parseResp.TokenType),
			)
			ctx = metadata.NewIncomingContext(ctx, metadata.Join(md, xTokenMD))
		}
		ctx = metadata.AppendToOutgoingContext(ctx,
			xTokenUserID, fmt.Sprintf("%v", parseResp.UserId),
			xTokenType, fmt.Sprintf("%d", parseResp.TokenType),
		)

		return ctx, nil
	}

}
