package handlers

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/helpers"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProduk(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var produk models.Produk

	// Parsing request body ke struct produk
	if err := c.BodyParser(&produk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Generate slug
	produk.Slug = helpers.GenerateSlug(produk.NamaProduk)

	// Ambil toko berdasarkan userID
	var toko models.Toko
	if err := database.DB.Where("id_user", userID).First(&toko).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Toko not found",
			"errors":  []string{"Toko not found"},
			"data":    nil,
		})
	}

	// Binding tokoID ke struct produk
	produk.IDToko = toko.ID

	// Simpan produk ke database
	if err := database.DB.Create(&produk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Preload relasi yang terkait: Toko dan Kategori
	if err := database.DB.Preload("Toko").Preload("Kategori").First(&produk, produk.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal memuat data Toko dan Kategori",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"data":    produk,
		"message": "Produk berhasil ditambahkan",
		"errors":  nil,
	})
}

func GetAllProduk(c *fiber.Ctx) error {
	var produk []models.Produk

	// Ambil seluruh data produk dari database
	if err := database.DB.Preload("Toko").Preload("Kategori").Find(&produk).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Produk not found",
			"errors":  []string{"Produk not found"},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success get all produk",
		"errors":  nil,
		"data":    produk,
	})
}
