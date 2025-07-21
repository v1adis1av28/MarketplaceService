package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsTokenExpired(tokenString string) (bool, error) {
	token, err := PrepareToken(tokenString)
	if err != nil {
		return false, fmt.Errorf("token parsing failed: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, err := claims.GetExpirationTime()
		if err != nil {
			return false, fmt.Errorf("could not get expiration time: %w", err)
		}

		if exp == nil {
			return false, fmt.Errorf("token has no expiration time")
		}

		return time.Now().After(exp.Time), nil
	}

	return false, fmt.Errorf("invalid token")
}

func GetSubFromToken(tokenString string) (string, error) {
	token, err := PrepareToken(tokenString)

	if err != nil {
		return "", fmt.Errorf("token parsing failed: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, err := claims.GetSubject()
		if err != nil {
			return "", fmt.Errorf("could not get sub: %w", err)
		}

		if email == "" {
			return "", fmt.Errorf("token has no sub time")
		}

		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}

func PrepareToken(tokenString string) (*jwt.Token, error) {
	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("jwtSecret"), nil
	})
	return token, err
}
