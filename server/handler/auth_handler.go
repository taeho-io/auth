package handler

import (
	"context"
	"time"

	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
)

type AuthHandlerFunc func(context.Context, *auth.AuthRequest) (*auth.AuthResponse, error)

func Auth(accessTokenExpiringDuration time.Duration, tkn token.Token) AuthHandlerFunc {
	return func(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
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
