package token

import "time"

type Config interface {
	SigningMethod() string
	SigningKey() string
	TokenIssuer() string
	AccessTokenExpireDuration() time.Duration
	RefreshTokenExpireDuration() time.Duration
}

type JWTConfig struct {
	Config

	signingMethod              string
	signingKey                 string
	tokenIssuer                string
	accessTokenExpireDuration  time.Duration
	refreshTokenExpireDuration time.Duration
}

func NewConfig(
	signingMethod string,
	signingKey string,
	tokenIssuer string,
	accessTokenExpireDuration time.Duration,
	refreshTokenExpireDuration time.Duration,
) Config {
	return &JWTConfig{
		signingMethod:              signingMethod,
		signingKey:                 signingKey,
		tokenIssuer:                tokenIssuer,
		accessTokenExpireDuration:  accessTokenExpireDuration,
		refreshTokenExpireDuration: refreshTokenExpireDuration,
	}
}

func MockConfig() Config {
	return NewConfig(
		"HS512",
		"MOCK_SIGNING_KEY",
		"MOCK_TOKEN_ISSUER",
		time.Hour,
		time.Hour*24*365,
	)
}

func (c *JWTConfig) SigningMethod() string {
	return c.signingMethod
}

func (c *JWTConfig) SigningKey() string {
	return c.signingKey
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
