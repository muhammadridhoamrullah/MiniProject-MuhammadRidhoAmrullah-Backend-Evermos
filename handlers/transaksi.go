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

func UpdateTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	transaksiID := c.Params("id")

	var inputUpdateTransaksi models.InputEditTransaksi
	// Parsing request body ke struct inputUpdateTransaksi
	if err := c.BodyParser(&inputUpdateTransaksi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Cari transaksi berdasarkan transaksiID
	var transaksi models.Transaksi
	if err := database.DB.Where("id = ?", transaksiID).First(&transaksi).Error; err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Transaksi not found",
			"errors":  []string{"Transaksi not found"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang memiliki transaksi
	if userID != transaksi.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to update this transaksi",
			"errors":  []string{"You have no access to update this transaksi"},
			"data":    nil,
		})
	}

	if inputUpdateTransaksi.AlamatPengiriman != "" {
		transaksi.AlamatPengiriman = inputUpdateTransaksi.AlamatPengiriman
	}

	if inputUpdateTransaksi.Status != "" {
		transaksi.Status = inputUpdateTransaksi.Status
	}

	if inputUpdateTransaksi.HargaTotal > 0 {
		transaksi.HargaTotal = inputUpdateTransaksi.HargaTotal
	}

	if inputUpdateTransaksi.MethodBayar != "" {
		transaksi.MethodBayar = inputUpdateTransaksi.MethodBayar
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&transaksi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to update transaksi",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success update transaksi",
		"errors":  nil,
		"data":    transaksi,
	})
}

func DeleteTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	transaksiID := c.Params("id")

	// Cari transaksi berdasarkan transaksiID
	var transaksi models.Transaksi
	if err := database.DB.Where("id = ?", transaksiID).First(&transaksi).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Transaksi not found",
			"errors":  []string{"Transaksi not found"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang memiliki transaksi
	if userID != transaksi.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to delete this transaksi",
			"errors":  []string{"You have no access to delete this transaksi"},
			"data":    nil,
		})
	}

	if err := database.DB.Unscoped().Delete(&transaksi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to delete transaksi",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success delete transaksi",
	})
}
