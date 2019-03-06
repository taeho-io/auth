package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/taeho-io/idl/gen/go/auth"
)

type Token interface {
	NewAccessToken(Claims) (string, error)
	NewRefreshToken(Claims) (string, error)
	VerifyToken(string) error
	ParseToken(string) (Claims, error)
}

type JWTToken struct {
	Token

	cfg          Config
	signingKey   *rsa.PrivateKey
	verifyingKey *rsa.PublicKey
}

func New(cfg Config) (Token, error) {
	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(cfg.SigningPEM()))
	if err != nil {
		return nil, err
	}
	verifyingKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cfg.VerifyingPEM()))
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		cfg:          cfg,
		signingKey:   signingKey,
		verifyingKey: verifyingKey,
	}, nil
}

func Mock() Token {
	tkn, _ := New(MockConfig())
	return tkn
}

func jwtSigningMethod(sm string) (*jwt.SigningMethodRSA, error) {
	switch sm {
	case "RS256":
		return jwt.SigningMethodRS256, nil
	case "RS512":
		return jwt.SigningMethodRS512, nil
	default:
		return nil, errors.New("invalid signing method")
	}
}

func (t *JWTToken) NewToken(expiresAt int64, claims Claims) (string, error) {
	signingMethod, err := jwtSigningMethod(t.cfg.SigningMethod())
	if err != nil {
		return "", err
	}

	verifiedClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.cfg.TokenIssuer(),
			IssuedAt:  time.Now().Unix(),
			Id:        xid.New().String(),
			ExpiresAt: expiresAt,
		},
		TokenType: claims.TokenType,
		UserID:    claims.UserID,
	}

	token := jwt.NewWithClaims(signingMethod, verifiedClaims)
	return token.SignedString(t.signingKey)
}

func (t *JWTToken) NewAccessToken(claims Claims) (accessToken string, err error) {
	claims.TokenType = auth.TokenType_ACCESS_TOKEN
	return t.NewToken(time.Now().Add(t.cfg.AccessTokenExpireDuration()).Unix(), claims)
}

func (t *JWTToken) NewRefreshToken(claims Claims) (refreshToken string, err error) {
	claims.TokenType = auth.TokenType_REFRESH_TOKEN
	return t.NewToken(time.Now().Add(t.cfg.RefreshTokenExpireDuration()).Unix(), claims)
}

func (t *JWTToken) verifyingKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return t.verifyingKey, nil
}

func (t *JWTToken) VerifyToken(token string) (err error) {
	_, err = jwt.Parse(token, t.verifyingKeyFunc)
	return
}

func (t *JWTToken) ParseToken(tokenString string) (claims Claims, err error) {
	claims = Claims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, t.verifyingKeyFunc)
	if err != nil {
		return
	}
	return
}
