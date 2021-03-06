package server

import (
	"crypto/md5"
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Config interface {
	Settings() Settings
	SigningKey() *rsa.PrivateKey
	VerifyingKey() *rsa.PublicKey
	KeyID() string
}

type DefaultConfig struct {
	Config

	settings     Settings
	signingKey   *rsa.PrivateKey
	verifyingKey *rsa.PublicKey
	keyID        string
}

func NewConfig(settings Settings) (Config, error) {
	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(settings.SigningPEM))
	if err != nil {
		return nil, err
	}
	verifyingKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(settings.VerifyingPEM))
	if err != nil {
		return nil, err
	}

	return &DefaultConfig{
		settings:     settings,
		signingKey:   signingKey,
		verifyingKey: verifyingKey,
		keyID:        fmt.Sprintf("%x", md5.Sum([]byte(settings.VerifyingPEM))),
	}, nil
}

func MockConfig() Config {
	cfg, _ := NewConfig(MockSettings())
	return cfg
}

func (c *DefaultConfig) Settings() Settings {
	return c.settings
}

func (c *DefaultConfig) SigningKey() *rsa.PrivateKey {
	return c.signingKey
}

func (c *DefaultConfig) VerifyingKey() *rsa.PublicKey {
	return c.verifyingKey
}

func (c *DefaultConfig) KeyID() string {
	return c.keyID
}
