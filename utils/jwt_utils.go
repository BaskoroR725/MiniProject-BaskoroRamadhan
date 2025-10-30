package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Struktur klaim JWT
type JWTClaim struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(getSecret())

// Ambil secret key dari .env
func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "evermos-secret-key" // fallback jika .env kosong
	}
	return secret
}

// Generate token baru
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return 0, errors.New("token tidak valid")
	}

	return claims.UserID, nil
}
