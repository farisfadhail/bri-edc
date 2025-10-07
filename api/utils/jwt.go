package utils

import (
	"bri-edc/api/config"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte
var jwtExpiration time.Duration

func InitJWT() {
	config.LoadEnv()

	jwtKey = []byte(config.GetEnv("JWT_SECRET", "JWT_SECRET"))

	expStr := config.GetEnv("JWT_EXPIRATION", "3600")
	expInt, err := strconv.Atoi(expStr)
	if err != nil {
		expInt = 3600
	}

	jwtExpiration = time.Duration(expInt) * time.Second
}

func GenerateJWT(terminalID string) (string, error) {
	claims := jwt.MapClaims{
		"terminal_id": terminalID,
		"exp":         time.Now().Add(jwtExpiration).Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token or claims")
	}

	terminalID, ok := claims["terminal_id"].(string)
	if !ok {
		return "", errors.New("terminal id not found in token claims")
	}

	return terminalID, nil
}
