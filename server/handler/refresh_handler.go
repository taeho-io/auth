package handler

import (
	"time"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidToken = status.Error(codes.Unauthenticated, "invalid token")
)

type RefreshHandlerFunc func(context.Context, *auth.RefreshRequest) (*auth.RefreshResponse, error)

func Refresh(accessTokenExpiringDuration time.Duration, tkn token.Token) RefreshHandlerFunc {
	return func(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
		claims, err := tkn.ParseToken(req.RefreshToken)
		if err != nil {
			return nil, ErrInvalidToken
		}

		if claims.TokenType != auth.TokenType_REFRESH_TOKEN {
			return nil, ErrInvalidToken
		}

		if err := req.Validate(); err != nil {
			return nil, ErrInvalidToken
		}

		accessTokenString, err := tkn.NewAccessToken(token.Claims{UserID: claims.UserID})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return &auth.RefreshResponse{
			AccessToken:  accessTokenString,
			ExpiresIn:    int64(accessTokenExpiringDuration.Seconds()),
			UserId:       claims.UserID,
			RefreshToken: req.RefreshToken,
		}, nil
	}
}
