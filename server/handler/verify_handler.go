package handler

import (
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VerifyHandlerFunc func(context.Context, *auth.VerifyRequest) (*auth.VerifyResponse, error)

func Verify(tkn token.Token) VerifyHandlerFunc {
	return func(ctx context.Context, req *auth.VerifyRequest) (*auth.VerifyResponse, error) {
		if err := req.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		err := tkn.VerifyToken(req.AccessToken)
		isValid := err == nil
		return &auth.VerifyResponse{
			IsValid: isValid,
		}, nil
	}
}
