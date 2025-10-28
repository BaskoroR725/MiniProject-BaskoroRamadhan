package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("evermos-secret-key") // ganti dengan key rahasia milikmu

// GenerateJWT membuat token JWT baru berdasarkan userID
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken memverifikasi token JWT dan mengembalikan userID jika valid
func ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		idFloat := claims["user_id"].(float64)
		return uint(idFloat), nil
	}

	return 0, err
}
