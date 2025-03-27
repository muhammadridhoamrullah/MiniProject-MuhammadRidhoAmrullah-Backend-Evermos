package middleware

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	// Cek user
	var user models.User
	if err := database.DB.Where("id", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User not found",
			"errors":  []string{"User not found"},
			"data":    nil,
		})
	}

	// Cek apakah user adalah admin
	if !user.IsAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Forbidden: You have no access",
			"errors":  []string{"Forbidden: You have no access"},
			"data":    nil,
		})
	}

	return c.Next()

}
