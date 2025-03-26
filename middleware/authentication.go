package middleware

import (
	"fmt"
	"mini-project-BE-Evermos/helpers"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(c *fiber.Ctx) error {

	// Ambil token dari header
	token := c.Get("Authorization")
	fmt.Println(token, "ini token di middleware authentication.go")

	// Cek jika token tidak ada, maka kembalikan response error
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Missing token",
			"errors":  []string{"Missing token"},
			"data":    nil,
		})
	}

	// Pisahkan token dari "Bearer"
	tokenString := strings.Split(token, " ")
	if len(tokenString) != 2 || tokenString[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid token",
			"errors":  []string{"Invalid token"},
			"data":    nil,
		})
	}

	fmt.Println(tokenString, "ini tokenString di middleware authentication.go")

	tokenNew := tokenString[1]

	fmt.Println(tokenNew, "ini tokenNew di middleware authentication.go")

	// Cek token apakah valid
	parsedToken, err := helpers.VerifyToken(tokenNew)
	if err != nil || !parsedToken.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid token or token expired",
			"errors":  []string{"Invalid token or token expired"},
			"data":    nil,
		})
	}

	fmt.Println(parsedToken, "ini parsedToken di middleware authentication.go")

	// Ambil user_id dari token
	claims := parsedToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	fmt.Println(claims, "ini claims di middleware authentication.go")
	fmt.Println(userID, "ini userID di middleware authentication.go")
	fmt.Printf("Tipe data userID: %T\n", userID)

	// Set user_id ke dalam context
	c.Locals("user_id", userID)

	return c.Next()
}
