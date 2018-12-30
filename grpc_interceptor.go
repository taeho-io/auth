package auth

import (
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	bearerScheme = "bearer"
	xTokenPrefix = "x-token-"
	xTokenUserID = "x-token-user_id"
)

func TokenUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	interface{},
	error,
) {
	return grpc_auth.UnaryServerInterceptor(authFunc)(ctx, req, info, handler)
}

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, bearerScheme)
	if err != nil {
		return ctx, nil
	}

	parseResp, err := GetAuthClient().Parse(ctx, &ParseRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
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

	// Set newly parsed x-token- metadata
	ctx = metadata.AppendToOutgoingContext(ctx, xTokenUserID, string(parseResp.UserId))

	ctx = ctxlogrus.ToContext(ctx, logrus.WithField(xTokenUserID, parseResp.UserId))

	return ctx, nil
}
