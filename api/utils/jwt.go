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

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(jwtExpiration).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("username not found in token claims")
	}

	return &username, nil
}
