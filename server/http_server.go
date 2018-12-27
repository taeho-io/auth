package server

import (
	"encoding/json"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
)

type JWKS struct {
	Keys []jwk.Key `json:"keys"`
}

func NewHttpServer(cfg Config) (*http.Server, error) {
	router := http.NewServeMux()
	router.HandleFunc("/jwks", func(w http.ResponseWriter, req *http.Request) {
		key, _ := jwk.New(cfg.VerifyingKey())

		_ = json.NewEncoder(w).Encode(&JWKS{Keys: []jwk.Key{key}})
	})

	httpServer := &http.Server{
		Addr:    ":81",
		Handler: router,
	}
	return httpServer, nil
}
