package server

import (
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
)

type Config interface {
	Settings() Settings
	SigningKey() *rsa.PrivateKey
	VerifyingKey() *rsa.PublicKey
}

type DefaultConfig struct {
	Config

	settings     Settings
	signingKey   *rsa.PrivateKey
	verifyingKey *rsa.PublicKey
}

func NewConfig(settings Settings) (Config, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(settings.SigningPEM))
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(settings.VerifyingPEM))
	if err != nil {
		return nil, err
	}

	return &DefaultConfig{
		settings:     settings,
		signingKey:   privateKey,
		verifyingKey: publicKey,
	}, nil
}

func MockConfig() Config {
	cfg, _ := NewConfig(MockSettings())
	return cfg
}

func (c *DefaultConfig) Settings() Settings {
	return c.settings
}
