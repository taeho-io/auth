package handler

import (
	"github.com/taeho-io/auth"
	"github.com/taeho-io/auth/pkg/token"
	"golang.org/x/net/context"
)

type ValidateHandlerFunc func(context.Context, *auth.ValidateRequest) (*auth.ValidateResponse, error)

func Validate(tkn token.Token) ValidateHandlerFunc {
	return func(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
		err := tkn.ValidateToken(req.AccessToken)
		isValid := err == nil
		return &auth.ValidateResponse{
			IsValid: isValid,
		}, nil
	}
}
