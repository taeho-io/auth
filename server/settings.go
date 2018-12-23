package server

import (
	"os"
	"time"
)

type Settings struct {
	SigningMethod                string
	SigningKey                   string
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
		SigningMethod:                getEnv("AUTH_SIGNING_METHOD", "HS512"),
		SigningKey:                   getEnv("AUTH_SIGNING_KEY", "DEFAULT_AUTH_SIGNING_KEY"),
		TokenIssuer:                  getEnv("AUTH_TOKEN_ISSUER", "DEFAULT_AUTH_TOKEN_ISSUER"),
		AccessTokenExpiringDuration:  time.Hour,
		RefreshTokenExpiringDuration: time.Hour * 24 * 365,
	}
}

func MockSettings() Settings {
	return Settings{
		SigningMethod:                "HS512",
		SigningKey:                   "MOCK_AUTH_SIGNING_KEY",
		TokenIssuer:                  "MOCK_AUTH_TOKEN_ISSUER",
		AccessTokenExpiringDuration:  time.Hour,
		RefreshTokenExpiringDuration: time.Hour * 24 * 365,
	}
}
