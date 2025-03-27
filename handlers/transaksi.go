package handlers

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaksi(c *fiber.Ctx) error {

	userID := c.Locals("user_id").(uint)

	var transaksi models.Transaksi
	// Parsing body request ke struct transaksi
	if err := c.BodyParser(&transaksi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Binding userID ke struct transaksi
	transaksi.IDUser = userID
	transaksi.Status = "PENDING"

	// Simpan transaksi ke database
	if err := database.DB.Create(&transaksi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   transaksi,
	})
}

func GetAllTransaksiByUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var transaksi []models.Transaksi

	// Mendapatkan seluruh transaksi dari user yang sedang login
	if err := database.DB.Preload("User").Where("id_user", userID).Find(&transaksi).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Transaksi tidak ditemukan",
			"errors":  []string{"Transaksi tidak ditemukan"},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   transaksi,
	})

}
