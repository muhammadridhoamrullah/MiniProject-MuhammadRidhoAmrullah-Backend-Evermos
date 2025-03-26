// Menandakan bahwa ini adalah entry point dari aplikasi.
// Wajib ada untuk file utama (main.go) yang akan dijalankan.
package main

import (
	"mini-project-BE-Evermos/database"
	"mini-project-BE-Evermos/routes"

	"github.com/gofiber/fiber/v2"
)

// Fungsi main yang akan dijalankan pertama kali saat aplikasi dijalankan
func main() {
	// Membuat instance dari Fiber
	app := fiber.New()

	// Menghubungkan ke database
	database.ConnectDB()

	// Memanggil Run Migration yang berfungsi untuk jika tabel belum di database ada maka akan membuat tabel
	database.RunMigration()

	// Menjalankan fungsi SetupRoutes
	routes.SetupRoutes(app)

	// Menjalankan aplikasi pada port 3000
	app.Listen(":3000")
}


