package encryption

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/techwithmat/kanbi/pkg/env"
)

type JWTCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)), // 7 days
	})

	signedToken, err := token.SignedString([]byte(env.MustGet("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", fmt.Errorf("create JWT token: %w", err)
	}

	return signedToken, nil
}

func VerifyAccessToken(accessToken string) (jwt.MapClaims, error) {
	// Validate the JWT
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(env.MustGet("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing JWT Token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
