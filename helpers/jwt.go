package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userId uint) (string, error) {
	// Buat claim payload
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(), 
		// exp adalah waktu expired token, dalam kasus ini 72 jam
	}

	// Buat token dengan HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ambil secret key dari .env
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))

}