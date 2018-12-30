package handler

import (
	"context"
	"time"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandlerFunc func(context.Context, *auth.AuthRequest) (*auth.AuthResponse, error)

func Auth(accessTokenExpiringDuration time.Duration, tkn token.Token) AuthHandlerFunc {
	return func(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		accessTokenString, err := tkn.NewAccessToken(token.Claims{UserID: req.UserId})
		if err != nil {
			return nil, err
		}

		refreshTokenString, err := tkn.NewRefreshToken(token.Claims{UserID: req.UserId})
		if err != nil {
			return nil, err
		}

		return &auth.AuthResponse{
			UserId:       req.UserId,
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString,
			ExpiresIn:    int64(accessTokenExpiringDuration.Seconds()),
		}, nil
	}

}
