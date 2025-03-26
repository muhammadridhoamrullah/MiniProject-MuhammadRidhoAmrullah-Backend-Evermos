package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userId uint) (string, error) {
	// Buat claim payload
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
		// exp adalah waktu expired token, dalam kasus ini 72 jam
	}

	// Buat token dengan HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ambil secret key dari .env
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Membaca isi dari token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Cek apakah token menggunakan metode HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// Cek jika terjadi error
	if err != nil {
		return nil, err
	}
	fmt.Println(token, "ini token di helpers jwt.go")
	return token, nil
}
