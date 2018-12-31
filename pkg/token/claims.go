package token

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/taeho-io/auth"
	"golang.org/x/net/context"
)

const (
	xTokenUserID = "x-token-user_id"
	xTokenType   = "x-token-type"
)

type Claims struct {
	jwt.StandardClaims

	TokenType auth.TokenType `json:"token_type"`
	UserID    int64          `json:"user_id"`
}

func GetClaimsFromMD(ctx context.Context) (*Claims, error) {
	md := metautils.ExtractIncoming(ctx)

	userIDString := md.Get(xTokenUserID)
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return nil, err
	}

	tokenTypeString := md.Get(xTokenType)
	tokenType, err := strconv.ParseInt(tokenTypeString, 10, 32)
	if err != nil {
		return nil, err
	}

	return &Claims{
		UserID:    userID,
		TokenType: auth.TokenType(tokenType),
	}, nil
}
