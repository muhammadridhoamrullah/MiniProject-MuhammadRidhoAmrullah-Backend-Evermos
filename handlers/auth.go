package handlers

import (
	"fmt"
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/helpers"
	"mini-project-BE-Evermos/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// Membuat variabel user dari struct models.User
	var user models.User

	// Parsing request body ke struct user
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
	
	if err := database.DB.Where("no_telp = ?", user.NoTelp).First((&checkPhone)).Error; err == nil {
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

	// Otomatis membuat toko untuk user yang baru terdaftar
	toko := models.Toko{
		IDUser: user.ID,
		NamaToko: "Toko " + user.Nama,
		UrlFoto: "https://via.placeholder.com/150",
	}

	if err := database.DB.Create(&toko).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message": "Failed to create toko",
			"errors": []string{err.Error()},
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"message": "User created successfully",
		"errors": nil,
		"data": fiber.Map{
			"user": user,
			"toko": toko,
		},
	})
}

func Login(c *fiber.Ctx) error {
	// Membuat variabel LoginUser dari struct models.LoginUser
	var LoginUser models.LoginUser

	// Parsing request body ke struct LoginUser
	if err := c.BodyParser(&LoginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"message": "Invalid request body",
			"errors": []string{err.Error()},
			"data": nil,
		})
	}

	// Cari user berdasarkan email
	var user models.User
	if err := database.DB.Where("email = ?", LoginUser.Email).First(&user).Error; err != nil {
		fmt.Println(err, "error disini")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"message": "User not found",
			"errors": []string{"Invalid Email or Password"},
			"data" : nil,
		})
	}

	// Cek password
	if !helpers.CheckPassword(user.KataSandi, LoginUser.KataSandi) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"message": "Invalid Email or Password",
			"errors": []string{"Invalid Email or Password"},
			"data": nil,
		})
	} 

	// Generate token menggunakan JWT
	token, err := helpers.GenerateJWT(user.ID) 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": false,
			"message" : "Failed to generate token",
			"errors": []string{err.Error()},
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"message": "Login successful",
		"errors" : nil,
		"data": fiber.Map{
			"token": token,
		},
	})

}