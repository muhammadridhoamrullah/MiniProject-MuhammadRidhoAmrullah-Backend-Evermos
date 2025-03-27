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

func UpdateAlamat(c *fiber.Ctx) error {
	// Ambil userID dari context
	userID := c.Locals("user_id").(uint)

	// Ambil ID alamat dari parameter URL
	alamatID := c.Params("id")

	var InputUpdateAlamat models.InputUpdateAlamat
	// Parsing request body ke struct alamat
	if err := c.BodyParser(&InputUpdateAlamat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Cari alamat berdasarkan ID dan userID
	var alamat models.Alamat

	if err := database.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&alamat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat tidak ditemukan",
			"errors":  []string{"Alamat tidak ditemukan"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang memiliki alamat
	if userID != alamat.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Anda tidak memiliki akses untuk memperbarui alamat ini",
			"errors":  []string{"Anda tidak memiliki akses untuk memperbarui alamat ini"},
			"data":    nil,
		})
	}

	if InputUpdateAlamat.JudulAlamat != "" {
		alamat.JudulAlamat = InputUpdateAlamat.JudulAlamat
	}

	if InputUpdateAlamat.NamaPenerima != "" {
		alamat.NamaPenerima = InputUpdateAlamat.NamaPenerima
	}

	if InputUpdateAlamat.NoTelp != "" {
		alamat.NoTelp = InputUpdateAlamat.NoTelp
	}

	if InputUpdateAlamat.DetailAlamat != "" {
		alamat.DetailAlamat = InputUpdateAlamat.DetailAlamat
	}

	if InputUpdateAlamat.IDProvinsi != "" {
		alamat.IDProvinsi = InputUpdateAlamat.IDProvinsi
	}

	if InputUpdateAlamat.IDKota != "" {
		alamat.IDKota = InputUpdateAlamat.IDKota
	}

	if InputUpdateAlamat.IDKecamatan != "" {
		alamat.IDKecamatan = InputUpdateAlamat.IDKecamatan
	}

	if InputUpdateAlamat.IDKelurahan != "" {
		alamat.IDKelurahan = InputUpdateAlamat.IDKelurahan
	}

	// Update alamat tergantung input yang dikirim
	if err := database.DB.Save(&alamat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal memperbarui alamat",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil memperbarui alamat",
	})
}

func DeleteAlamat(c *fiber.Ctx) error {
	// Ambil userID dari context
	userID := c.Locals("user_id").(uint)

	// Ambil ID alamat dari parameter URL
	alamatID := c.Params("id")

	// Cari alamat berdasarkan ID
	var alamat models.Alamat

	if err := database.DB.Where("id = ? ", alamatID).First(&alamat).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat tidak ditemukan",
			"errors":  []string{"Alamat tidak ditemukan"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang memiliki alamat
	if userID != alamat.IDUser {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Anda tidak memiliki akses untuk menghapus alamat ini",
			"errors":  []string{"Anda tidak memiliki akses untuk menghapus alamat ini"},
			"data":    nil,
		})
	}

	if err := database.DB.Unscoped().Delete(&alamat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal menghapus alamat",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Berhasil menghapus alamat",
	})
}
