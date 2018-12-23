package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type Token interface {
	NewAccessToken(Claims) (string, error)
	NewRefreshToken(Claims) (string, error)
	ValidateToken(string) error
	ParseToken(string) (Claims, error)
}

type JWTToken struct {
	Token

	cfg Config
}

func New(cfg Config) Token {
	return &JWTToken{
		cfg: cfg,
	}
}

func Mock() Token {
	return New(MockConfig())
}

func jwtSigningMethod(sm string) (*jwt.SigningMethodHMAC, error) {
	switch sm {
	case "HS256":
		return jwt.SigningMethodHS256, nil
	case "HS512":
		return jwt.SigningMethodHS512, nil
	default:
		return nil, errors.New("invalid signing method")
	}
}

func (t *JWTToken) NewAccessToken(claims Claims) (accessToken string, err error) {
	signingMethod, err := jwtSigningMethod(t.cfg.SigningMethod())
	if err != nil {
		return "", err
	}

	verifiedClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.cfg.TokenIssuer(),
			IssuedAt:  time.Now().Unix(),
			Id:        xid.New().String(),
			ExpiresAt: time.Now().Add(t.cfg.AccessTokenExpireDuration()).Unix(),
		},
		UserID: claims.UserID,
	}

	token := jwt.NewWithClaims(signingMethod, verifiedClaims)
	return token.SignedString([]byte(t.cfg.SigningKey()))
}

func (t *JWTToken) NewRefreshToken(claims Claims) (refreshToken string, err error) {
	signingMethod, err := jwtSigningMethod(t.cfg.SigningMethod())
	if err != nil {
		return "", err
	}

	verifiedClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    t.cfg.TokenIssuer(),
			IssuedAt:  time.Now().Unix(),
			Id:        xid.New().String(),
			ExpiresAt: time.Now().Add(t.cfg.RefreshTokenExpireDuration()).Unix(),
		},
		UserID: claims.UserID,
	}

	token := jwt.NewWithClaims(signingMethod, verifiedClaims)
	return token.SignedString([]byte(t.cfg.SigningKey()))
}

func (t *JWTToken) signingKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(t.cfg.SigningKey()), nil
}

func (t *JWTToken) ValidateToken(token string) (err error) {
	_, err = jwt.Parse(token, t.signingKeyFunc)
	return
}

func (t *JWTToken) ParseToken(tokenString string) (claims Claims, err error) {
	claims = Claims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, t.signingKeyFunc)
	if err != nil {
		return
	}
	return
}
