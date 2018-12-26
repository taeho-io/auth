package server

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

type JWKS struct {
	Keys []jwk.Key `json:"keys"`
}

func NewHttpServer(cfg Config) (*http.Server, error) {
	router := http.NewServeMux()
	router.HandleFunc("/jwks", func(w http.ResponseWriter, req *http.Request) {
		publicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cfg.Settings().VerifyingPEM))

		key, _ := jwk.New(publicKey)
		_ = key.Set("alg", cfg.Settings().SigningMethod)
		_ = key.Set("use", "sig")

		_ = json.NewEncoder(w).Encode(&JWKS{Keys: []jwk.Key{key}})
	})

	httpServer := &http.Server{
		Addr:    ":80",
		Handler: router,
	}
	return httpServer, nil
}
