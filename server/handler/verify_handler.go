package handler

import (
	"github.com/taeho-io/auth/pkg/token"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
)

type VerifyHandlerFunc func(context.Context, *auth.VerifyRequest) (*auth.VerifyResponse, error)

func Verify(tkn token.Token) VerifyHandlerFunc {
	return func(ctx context.Context, req *auth.VerifyRequest) (*auth.VerifyResponse, error) {
		if err := req.Validate(); err != nil {
			return &auth.VerifyResponse{
				IsValid: false,
			}, nil
		}

		err := tkn.VerifyToken(req.AccessToken)
		isValid := err == nil
		return &auth.VerifyResponse{
			IsValid: isValid,
		}, nil
	}
}
