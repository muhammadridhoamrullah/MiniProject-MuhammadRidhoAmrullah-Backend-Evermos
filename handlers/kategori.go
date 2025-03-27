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

func UpdateKategori(c *fiber.Ctx) error {
	var inputEditKategori models.InputUpdateKategori

	// Parsing request body ke struct kategori
	if err := c.BodyParser(&inputEditKategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	id := c.Params("id")

	var kategori models.Kategori

	// Cari kategori berdasarkan ID
	if err := database.DB.Where("id = ?", id).First(&kategori).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Kategori not found",
			"errors":  []string{"Kategori not found"},
			"data":    nil,
		})
	}

	// Inputan yang di update
	if inputEditKategori.NamaKategori != "" {
		kategori.NamaKategori = inputEditKategori.NamaKategori
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&kategori).Error; err != nil {
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

func DeleteKategori(c *fiber.Ctx) error {
	id := c.Params("id")

	var kategori models.Kategori

	// Cari kategori berdasarkan ID
	if err := database.DB.Where("id = ?", id).First(&kategori).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Kategori not found",
			"errors":  []string{"Kategori not found"},
			"data":    nil,
		})
	}

	// Hapus kategori dari database
	if err := database.DB.Unscoped().Delete(&kategori).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success delete kategori",
	})
}
