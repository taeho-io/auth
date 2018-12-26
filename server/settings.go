package server

import (
	"os"
	"time"

	"github.com/taeho-io/auth/pkg/token"
)

type Settings struct {
	SigningMethod                string
	SigningPEM                   string
	VerifyingPEM                 string
	TokenIssuer                  string
	AccessTokenExpiringDuration  time.Duration
	RefreshTokenExpiringDuration time.Duration
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewSettings() Settings {
	return Settings{
		SigningMethod:                getEnv("AUTH_SIGNING_METHOD", "RS512"),
		SigningPEM:                   getEnv("AUTH_SIGNING_PEM", token.MockSigningPEM),
		VerifyingPEM:                 getEnv("AUTH_VERIFYING_PEM", token.MockVerifyPEM),
		TokenIssuer:                  getEnv("AUTH_TOKEN_ISSUER", "DEFAULT_AUTH_TOKEN_ISSUER"),
		AccessTokenExpiringDuration:  time.Hour,
		RefreshTokenExpiringDuration: time.Hour * 24 * 365,
	}
}

func MockSettings() Settings {
	return Settings{
		SigningMethod:                "RS512",
		SigningPEM:                   token.MockSigningPEM,
		VerifyingPEM:                 token.MockVerifyPEM,
		TokenIssuer:                  "MOCK_AUTH_TOKEN_ISSUER",
		AccessTokenExpiringDuration:  time.Hour,
		RefreshTokenExpiringDuration: time.Hour * 24 * 365,
	}
}
