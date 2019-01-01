package handler

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/taeho-io/auth"
	"golang.org/x/net/context"
)

type JWKSHandlerFunc func(context.Context, *auth.JWKSRequest) (*auth.JWKSResponse, error)

func JWKS(verifyingKey *rsa.PublicKey) JWKSHandlerFunc {
	return func(ctx context.Context, response *auth.JWKSRequest) (*auth.JWKSResponse, error) {
		key, err := jwk.New(verifyingKey)
		if err != nil {
			return nil, err
		}

		keyMap := make(map[string]interface{})
		if err := key.PopulateMap(keyMap); err != nil {
			return nil, err
		}

		jwkMsg := &auth.JWK{
			Kty: key.KeyType().String(),
			E:   keyMap[`e`].(string),
			N:   keyMap[`n`].(string),
		}

		return &auth.JWKSResponse{
			Keys: []*auth.JWK{jwkMsg},
		}, nil
	}
}
