package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var SECRET_KEY = []byte("LIGHT_GIVE_ME_STRENGTH")

const ISSUER_CLAIM_VALUE = "gql.raidcomp.io"
const AUDIENCE_CLAIM_VALUE = "raidcomp.io"
const TOKEN_DURATION = time.Hour * 24

func ParseToken(tokenStr string) (string, time.Time, error) {
	claims := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		log.Fatalf("can not parse/verify token %v", err)
		return "", time.Time{}, err
	}

	return claims.ID, claims.ExpiresAt.Time, nil
}

func GenerateToken(ctx context.Context, userID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:   ISSUER_CLAIM_VALUE,
		Audience: []string{AUDIENCE_CLAIM_VALUE},
		// TODO: What to put here!??!
		Subject:   "",
		ExpiresAt: jwt.NewNumericDate(now.Add(TOKEN_DURATION)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		// TODO: maybe this is supposed to be a new UUID?
		ID: userID,
	})

	// TODO: Integrate with KMS
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Fatalf("Error in Generating key %v", err)
		return "", err
	}

	return tokenString, nil
}
