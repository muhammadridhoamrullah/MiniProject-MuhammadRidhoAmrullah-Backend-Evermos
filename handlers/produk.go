package handlers

import (
	"fmt"
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

func UpdateProduk(c *fiber.Ctx) error {
	produkID := c.Params("id")
	userID := c.Locals("user_id").(uint)

	var InputUpdateProduk models.InputUpdateProduk
	// Parsing request body ke struct produk
	if err := c.BodyParser(&InputUpdateProduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Cari produk berdasarkan ID
	var produk models.Produk

	if err := database.DB.Preload("Toko.User").Where("id = ?", produkID).First(&produk).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Produk not found",
			"errors":  []string{"Produk not found"},
			"data":    nil,
		})
	}

	produk.IDToko = userID

	// Cek apakah user yang login sama dengan user yang memiliki produk
	fmt.Println("userID", userID)
	fmt.Println("produk.Toko.IDUser", produk.Toko.IDUser)
	if userID != produk.Toko.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to update other user's produk",
			"errors":  []string{"You have no access to update other user's produk"},
			"data":    nil,
		})
	}

	// Update produk tergantung input yang dikirim
	if InputUpdateProduk.NamaProduk != "" {
		produk.NamaProduk = InputUpdateProduk.NamaProduk
		produk.Slug = helpers.GenerateSlug(InputUpdateProduk.NamaProduk)
	}

	if InputUpdateProduk.HargaReseller != "" {
		produk.HargaReseller = InputUpdateProduk.HargaReseller
	}

	if InputUpdateProduk.HargaKonsumen != "" {
		produk.HargaKonsumen = InputUpdateProduk.HargaKonsumen
	}

	if InputUpdateProduk.Stok >= 0 {
		produk.Stok = InputUpdateProduk.Stok
	}

	if InputUpdateProduk.Deskripsi != "" {
		produk.Deskripsi = InputUpdateProduk.Deskripsi
	}

	if err := database.DB.Save(&produk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to update produk",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success update produk",
	})
}

func DeleteProduk(c *fiber.Ctx) error {
	produkID := c.Params("id")
	userID := c.Locals("user_id").(uint)

	// Cari produk berdasarkan ID
	var produk models.Produk

	if err := database.DB.Preload("Toko.User").Where("id = ?", produkID).First(&produk).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Produk not found",
			"errors":  []string{"Produk not found"},
			"data":    nil,
		})
	}

	fmt.Println("userID", userID)
	fmt.Println("produk.")
	// Cek apakah user yang login sama dengan user yang memiliki produk
	if userID != produk.Toko.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to delete other user's produk",
			"errors":  []string{"You have no access to delete other user's produk"},
			"data":    nil,
		})
	}

	if err := database.DB.Unscoped().Delete(&produk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to delete produk",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success delete produk",
	})
}
