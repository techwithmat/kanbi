package encryption

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/techwithmat/kanban-app/pkg/env"
)

func GenerateAccessToken(userId int, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userId
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix() // 7 days

	secret := []byte(env.MustGet("ACCESS_TOKEN_SECRET"))
	signedToken, err := token.SignedString(secret)

	if err != nil {
		return "", fmt.Errorf("create JWT token: %w", err)
	}

	return signedToken, nil
}
