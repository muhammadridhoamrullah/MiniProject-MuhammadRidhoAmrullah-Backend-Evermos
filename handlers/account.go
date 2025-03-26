package handlers

import (
	"fmt"
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	// Ambil user_id dari context
	userIDRaw := c.Locals("user_id")
	fmt.Printf("Tipe data userIDRaw di handlers: %T\n", userIDRaw)

	userID, ok := userIDRaw.(uint)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid userID",
			"errors":  []string{"Invalid userID"},
			"data":    nil,
		})
	}

	fmt.Println(userID, "ini userID di handlers handlers account.go")

	var user models.User
	// Cari user berdasarkan userID beserta relasi toko
	if err := database.DB.Preload("Toko").First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User not found",
			"errors":  []string{"User not found"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang akan diambil datanya
	if userID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to get other user's profile",
			"errors":  []string{"You have no access to get other user's profile"},
			"data":    nil,
		})
	}

	// Sembunyikan kata_sandi
	user.KataSandi = ""

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success get profile",
		"errors":  nil,
		"data":    user,
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	// Parsing request body ke struct inputEditProfile
	var inputEditProfile models.UpdateProfileUser
	if err := c.BodyParser(&inputEditProfile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request body",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Cari user berdasarkan userID
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User not found",
			"errors":  []string{"User not found"},
			"data":    nil,
		})
	}

	// Cek apakah user yang login sama dengan user yang akan diupdate
	if userID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You have no access to update other user's profile",
			"errors":  []string{"You have no access to update other user's profile"},
			"data":    nil,
		})
	}

	// Update user tergantung inputEditProfile
	updates := make(map[string]interface{})

	if inputEditProfile.Nama != "" {
		updates["Nama"] = inputEditProfile.Nama
	}
	if inputEditProfile.NoTelp != "" {
		updates["NoTelp"] = inputEditProfile.NoTelp
	}
	if inputEditProfile.TanggalLahir != "" {
		updates["TanggalLahir"] = inputEditProfile.TanggalLahir
	}
	if inputEditProfile.JenisKelamin != "" {
		updates["JenisKelamin"] = inputEditProfile.JenisKelamin
	}
	if inputEditProfile.Tentang != "" {
		updates["Tentang"] = inputEditProfile.Tentang
	}
	if inputEditProfile.Pekerjaan != "" {
		updates["Pekerjaan"] = inputEditProfile.Pekerjaan
	}
	if inputEditProfile.IDProvinsi != "" {
		updates["IDProvinsi"] = inputEditProfile.IDProvinsi
	}
	if inputEditProfile.IDKota != "" {
		updates["IDKota"] = inputEditProfile.IDKota
	}

	// Simpan perubahan
	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to update profile",
			"errors":  []string{err.Error()},
			"data":    nil,
		})
	}

	// Kembalikan response

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Success update profile",
		"errors":  nil,
	})
}
