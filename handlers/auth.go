package handlers

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/helpers"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser((&user)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message" : "Invalid request body",
			"errors": []string{err.Error()},
			"data": nil,
		})
	}

	// Cek email apakah exists
	var checkEmail models.User
	if err := database.DB.Where("email = ?", user.Email).First((&checkEmail)).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON((fiber.Map{
			"status": false,
			"message": "Email already exists",
			"errors" : []string{"Email already exists"},
			"data": nil,
		}))
	}

	// Cek nomor telepon apakah exists
	var checkPhone models.User
	
	if err := database.DB.Where("no_telp = ?", user.).First((&checkPhone)).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status" : false,
			"message": "Phone number already exists",
			"errors": []string{"Phone number already exists"},
			"data": nil,
		})
}


	// Hashing Password
	hashingPassword, err := helpers.HashPassword((user.KataSandi))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"message": "Failed to hash password",
			"errors": []string{err.Error()},
			"data": nil,
		})
	}
	user.KataSandi = string(hashingPassword)

	// Simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message": "Failed to create user",
			"errors": []string{err.Error()},
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"message": "User created successfully",
		"errors": nil,
		"data": user,
	})



}