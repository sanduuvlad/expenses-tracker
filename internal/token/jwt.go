package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenDuration = 24 * time.Hour

type TokenManager struct {
	secret string
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{
		secret: secret,
	}
}

func (tm *TokenManager) GenerateToken(userID int64) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,

		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenDuration)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenStr, err := token.SignedString([]byte(tm.secret))
	if err != nil {
		return "", fmt.Errorf("sign JWT token: %w", err)
	}

	return tokenStr, nil
}

func (tm *TokenManager) ValidateToken(tokenStr string) (*Claims, error) {
	var claims Claims

	tokenParse, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("invalid signing method: %v", token.Method)
			}

			return []byte(tm.secret), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("parse JWT token: %w", err)
	}

	if !tokenParse.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}

	return &claims, nil
}
