package token_service

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func CheckToken(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}

	token := strings.TrimPrefix(md["authorization"][0], "Bearer ")

	uc, err := parseAccessToken(token)

	if err != nil {
		return nil, errInvalidToken
	}

	return uc, nil
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	uc, err := parseAccessToken(token)

	if err != nil {
		log.Print(err)
		return false
	}

	if uc.StandardClaims.Valid() != nil {
		return false
	}

	return true
}

func parseAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if parsedAccessToken == nil {
		return nil, errInvalidToken
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}

// func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
//  refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

//  return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
// }

// func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
//  parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//   return []byte(os.Getenv("TOKEN_SECRET")), nil
//  })

//  return parsedRefreshToken.Claims.(*jwt.StandardClaims)
// }
