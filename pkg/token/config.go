package token

import (
	"time"
)

type Config interface {
	SigningMethod() string
	SigningPEM() string
	VerifyingPEM() string
	TokenIssuer() string
	AccessTokenExpireDuration() time.Duration
	RefreshTokenExpireDuration() time.Duration
}

type JWTConfig struct {
	Config

	signingMethod              string
	signingPEM                 string
	verifyingPEM               string
	tokenIssuer                string
	accessTokenExpireDuration  time.Duration
	refreshTokenExpireDuration time.Duration
}

func NewConfig(
	signingMethod string,
	signingPEM string,
	verifyingPEM string,
	tokenIssuer string,
	accessTokenExpireDuration time.Duration,
	refreshTokenExpireDuration time.Duration,
) Config {
	return &JWTConfig{
		signingMethod:              signingMethod,
		signingPEM:                 signingPEM,
		verifyingPEM:               verifyingPEM,
		tokenIssuer:                tokenIssuer,
		accessTokenExpireDuration:  accessTokenExpireDuration,
		refreshTokenExpireDuration: refreshTokenExpireDuration,
	}
}

func MockConfig() Config {
	return NewConfig(
		"RS512",
		MockSigningPEM,
		MockVerifyPEM,
		"MOCK_TOKEN_ISSUER",
		time.Hour,
		time.Hour*24*365,
	)
}

func (c *JWTConfig) SigningMethod() string {
	return c.signingMethod
}

func (c *JWTConfig) SigningPEM() string {
	return c.signingPEM
}

func (c *JWTConfig) VerifyingPEM() string {
	return c.verifyingPEM
}

func (c *JWTConfig) TokenIssuer() string {
	return c.tokenIssuer
}

func (c *JWTConfig) AccessTokenExpireDuration() time.Duration {
	return c.accessTokenExpireDuration
}

func (c *JWTConfig) RefreshTokenExpireDuration() time.Duration {
	return c.refreshTokenExpireDuration
}
