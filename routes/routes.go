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
	account.Delete("/delete", handlers.DeleteAccount)

	// account.Put("/password", handlers.ChangePassword)

	toko := app.Group("/toko", middleware.Authentication)

	toko.Get("/", handlers.GetToko)
	toko.Get("/all", handlers.GetAllToko)
	toko.Get("/search", handlers.SearchToko)
	toko.Put("/", handlers.UpdateToko)
	toko.Delete("/", handlers.DeleteToko)

	alamat := app.Group("/alamat", middleware.Authentication)
	alamat.Get("/", handlers.GetAllAlamatUser)   // Get semua alamat user
	alamat.Post("/", handlers.CreateAlamat)      // Tambah alamat baru
	alamat.Put("/:id", handlers.UpdateAlamat)    // Update alamat
	alamat.Delete("/:id", handlers.DeleteAlamat) // Hapus alamat

	// // Group untuk wilayah
	wilayah := app.Group("/wilayah")
	wilayah.Get("/provinsi", handlers.GetProvinces)                // Mendapatkan daftar provinsi
	wilayah.Get("/kota/:provinsiID", handlers.GetKotaByProvinsiId) // Mendapatkan daftar kota berdasarkan provinsi\
	wilayah.Get("/kecamatan/:kotaID", handlers.GetKecamatanByKotaID)
	wilayah.Get("/kelurahan/:kecamatanID", handlers.GetKelurahanByKecamatanID)

	kategori := app.Group("/kategori", middleware.Authentication, middleware.IsAdmin)
	kategori.Post("/", handlers.CreateKategori)
	kategori.Get("/", handlers.GetAllKategori)
	kategori.Put("/:id", handlers.UpdateKategori)
	kategori.Delete("/:id", handlers.DeleteKategori)
	// kategori.Get("/:id", handlers.GetKategoriById)

	produk := app.Group("/produk", middleware.Authentication)
	produk.Post("/", handlers.CreateProduk)
	produk.Get("/", handlers.GetAllProduk)
	produk.Put("/:id", handlers.UpdateProduk)
	produk.Delete("/:id", handlers.DeleteProduk)
	// produk.Get("/:id", handlers.GetProdukById)
	// produk.Get("/search", handlers.SearchProduk)
	// produk.Get("/filter", handlers.FilterProduk)
	// produk.Get("/kategori/:id", handlers.GetProdukByKategori)
	// produk.Get("/toko/:id", handlers.GetProdukByToko)

	transaksi := app.Group("/transaksi", middleware.Authentication)
	transaksi.Post("/", handlers.CreateTransaksi)
	transaksi.Get("/", handlers.GetAllTransaksiByUser)
	transaksi.Put("/:id", handlers.UpdateTransaksi)
	transaksi.Delete("/:id", handlers.DeleteTransaksi)
	// transaksi.Get("/:id", handlers.GetTransaksiById)
	// transaksi.Get("/toko/:id", handlers.GetTransaksiByToko)
	// transaksi.Get("/user/:id", handlers.GetTransaksiByUser)
	// transaksi.Get("/status/:status", handlers.GetTransaksiByStatus)

}
