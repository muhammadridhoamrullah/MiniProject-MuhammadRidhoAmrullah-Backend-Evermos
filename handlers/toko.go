package handlers

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var toko models.Toko
	// Cari toko berdasarkan userID
	if err := database.DB.Where("id_user", userID).First(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko not found",
			"errors":  []string{"Toko not found"},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success get toko",
		"errors":  nil,
		"data":    toko,
	})
}

func GetAllToko(c *fiber.Ctx) error {
	var toko []models.Toko
	// Cari semua toko
	if err := database.DB.Find(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko not found",
			"errors":  []string{"Toko not found"},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success get all toko",
		"errors":  nil,
		"data":    toko,
	})
}

func SearchToko(c *fiber.Ctx) error {
	// Ambil parameter query string untuk pencarian
	namaToko := c.Query("nama_toko", "") // Nama toko yang ingin dicari
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	// Parse page dan limit menjadi integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	// Hitung offset untuk pagination
	offset := (page - 1) * limit

	var toko []models.Toko
	query := database.DB.Model(&models.Toko{})

	// Filtering berdasarkan nama toko
	if namaToko != "" {
		query = query.Where("nama_toko LIKE ?", "%"+namaToko+"%")
	}

	// Menambahkan pagination dengan limit dan offset
	query = query.Offset(offset).Limit(limit)

	// Preload User untuk mendapatkan informasi pemilik toko
	if err := query.Preload("User").Find(&toko).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"error":  "Failed to search toko",
		})
	}

	// Mengembalikan hasil pencarian dengan informasi pagination
	return c.JSON(fiber.Map{
		"status": true,
		"data":   toko,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
		},
	})
}

func UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	// Parsing request body ke struct UpdateToko
	var inputEditToko models.UpdateToko
	if err := c.BodyParser(&inputEditToko); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Cari toko berdasarkan userID
	var toko models.Toko
	if err := database.DB.Where("id_user", userID).First(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko not found",
			"errors":  []string{"Toko not found"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang memiliki toko
	if userID != toko.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to update other user's toko",
			"errors":  []string{"You have no access to update other user's toko"},
			"data":    nil,
		})
	}

	// Update toko tergantung input yang dikirim

	if inputEditToko.NamaToko != "" {
		toko.NamaToko = inputEditToko.NamaToko
	}

	if inputEditToko.UrlFoto != "" {
		toko.UrlFoto = inputEditToko.UrlFoto
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&toko).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"error":   "Failed to update toko",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success update toko",
		"errors":  nil,
		"data":    toko,
	})

}

func DeleteToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	// Cari toko berdasarkan userID
	var toko models.Toko
	if err := database.DB.Where("id_user", userID).First(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko not found",
			"errors":  []string{"Toko not found"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang memiliki toko
	if userID != toko.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to delete other user's toko",
			"errors":  []string{"You have no access to delete other user's toko"},
			"data":    nil,
		})
	}

	// Hapus toko dari database
	if err := database.DB.Unscoped().Delete(&toko).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to delete toko",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success delete toko",
	})
}
