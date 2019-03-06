package handler

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
)

type JwksHandlerFunc func(context.Context, *auth.JwksRequest) (*auth.JwksResponse, error)

func Jwks(verifyingKey *rsa.PublicKey) JwksHandlerFunc {
	return func(ctx context.Context, response *auth.JwksRequest) (*auth.JwksResponse, error) {
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

		return &auth.JwksResponse{
			Keys: []*auth.JWK{jwkMsg},
		}, nil
	}
}
