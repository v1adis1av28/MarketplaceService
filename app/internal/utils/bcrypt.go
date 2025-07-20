package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
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

func IsTokenExpired(s string) (bool, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return "jwtSecret", nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expVal, ok := claims["exp"].(float64)
		if !ok {
			return false, fmt.Errorf("exp claim not found or invalid")
		}

		expTime := time.Unix(int64(expVal), 0)
		if time.Now().After(expTime) {
			return true, nil
		}
		return false, nil
	}
	return false, fmt.Errorf("token was expired")
}
