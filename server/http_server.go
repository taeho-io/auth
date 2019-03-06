package server

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lestrrat-go/jwx/jwk"
	authClient "github.com/taeho-io/auth"
	"github.com/taeho-io/idl/gen/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type JWKS struct {
	Keys []jwk.Key `json:"keys"`
}

func ServeHTTP(addr string, _ Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := auth.RegisterAuthHandlerFromEndpoint(
		ctx,
		mux,
		authClient.ServiceURL,
		opts,
	); err != nil {
		return err
	}

	return http.ListenAndServe(addr, mux)
}
