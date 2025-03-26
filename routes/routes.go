package routes

import (
	"mini-project-BE-Evermos/handlers"
	"mini-project-BE-Evermos/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// api := app.Group("/api")
	auth := app.Group("/")

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	account := app.Group("/account", middleware.Authentication)

	account.Get("/profile", handlers.GetProfile)
	account.Put("/profile", handlers.UpdateProfile)
	// account.Put("/password", handlers.ChangePassword)
	// account.Delete("/delete", handlers.DeleteAccount)

	toko := app.Group("/toko", middleware.Authentication)

	toko.Get("/", handlers.GetToko)
	toko.Get("/all", handlers.GetAllToko)
	toko.Get("/search", handlers.SearchToko)
	toko.Put("/", handlers.UpdateToko)
	// toko.Delete("/", handlers.DeleteToko)

	// alamat := app.Group("/alamat", middleware.Authentication)
	// alamat.Get("/", handlers.GetAllAlamat)        // Get semua alamat user
	// alamat.Get("/:id", handlers.GetAlamatByID)    // Get alamat berdasarkan ID
	// alamat.Post("/", handlers.CreateAlamat)       // Tambah alamat baru
	// alamat.Put("/:id", handlers.UpdateAlamat)     // Update alamat
	// alamat.Delete("/:id", handlers.DeleteAlamat)  // Hapus alamat
	// alamat.Put("/set-utama/:id", handlers.SetUtamaAlamat) // Set alamat utama
	
	// // Group untuk wilayah
	// wilayah := app.Group("/wilayah")
	// wilayah.Get("/provinces", handlers.GetProvinces)        // Mendapatkan daftar provinsi
	// wilayah.Get("/cities/:provinceID", handlers.GetCitiesByProvince)  // Mendapatkan daftar kota berdasarkan provinsi

}
