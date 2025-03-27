package handlers

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func CreateKategori(c *fiber.Ctx) error {
	var kategori models.Kategori

	// Parsing request body ke struct kategori
	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Simpan kategori ke database
	if err := database.DB.Create(&kategori).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   kategori,
	})
}

func GetAllKategori(c *fiber.Ctx) error {
	var kategori []models.Kategori

	// Ambil seluruh data kategori dari database
	if err := database.DB.Find(&kategori).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Kategori not found",
			"errors":  []string{"Kategori not found"},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success get all kategori",
		"errors":  nil,
		"data":    kategori,
	})

}
