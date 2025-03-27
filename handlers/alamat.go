package handlers

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllAlamatUser(c *fiber.Ctx) error {
	// Ambil userID dari context
	userID := c.Locals("user_id").(uint)

	var alamat []models.Alamat

	// Mendapatkan seluruh alamat dari user yang sedang login
	if err := database.DB.Preload("User").Where("id_user", userID).Find(&alamat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat tidak ditemukan",
			"errors":  []string{"Alamat tidak ditemukan"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": alamat,
	})
}

func CreateAlamat(c *fiber.Ctx) error {
	// Ambil userID dari context
	userID := c.Locals("user_id").(uint)

	var alamat models.Alamat

	// Parsing request body ke struct alamat
	if err := c.BodyParser(&alamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Binding userID ke struct alamat
	alamat.IDUser = userID

	// Simpan alamat ke database
	if err := database.DB.Create(&alamat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Preload relasi yang terkait: Provinsi, Kota, Kecamatan, Kelurahan
	if err := database.DB.Preload("Provinsi").Preload("Kota").Preload("Kecamatan").Preload("Kelurahan").First(&alamat, alamat.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal preload data wilayah",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil menambahkan alamat baru",
		"errors":  nil,
		"data":    alamat,
	})
}
