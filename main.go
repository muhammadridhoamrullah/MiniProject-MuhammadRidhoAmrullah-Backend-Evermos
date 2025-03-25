// Menandakan bahwa ini adalah entry point dari aplikasi.
// Wajib ada untuk file utama (main.go) yang akan dijalankan.
package main

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/models"
	"mini-project-BE-Evermos/routes"

	"github.com/gofiber/fiber/v2"
)

// Fungsi main yang akan dijalankan pertama kali saat aplikasi dijalankan
func main() {
	// Membuat instance dari Fiber
	app := fiber.New()

	// Menghubungkan ke database
	database.ConnectDB()

	// Membuat tabel user jika belum ada
	database.DB.AutoMigrate(&models.User{})

	// Menjalankan fungsi SetupRoutes
	routes.SetupRoutes(app)




	// Menjalankan aplikasi pada port 3000
	app.Listen(":3000")
}


