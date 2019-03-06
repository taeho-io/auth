package handler

import (
	"github.com/taeho-io/auth/pkg/token"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ParseHandlerFunc func(context.Context, *auth.ParseRequest) (*auth.ParseResponse, error)

func Parse(tkn token.Token) ParseHandlerFunc {
	return func(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		claims, err := tkn.ParseToken(req.AccessToken)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "Unauthorized")
		}

		return &auth.ParseResponse{
			UserId:    claims.UserID,
			TokenType: claims.TokenType,
		}, nil
	}
}
