package handlers

import (
	"encoding/json"
	"io"
	"mini-project-BE-Evermos/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetProvinces(c *fiber.Ctx) error {

	// Ambil data provinsi dari 3rd party api
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to get provinces",
			"errors":  []string{"Failed to get provinces"},
			"data":    nil,
		})
	}
	defer resp.Body.Close()

	// Membaca body response dari 3rd party api
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membaca body response provinsi",
		})
	}

	// Deklarasi variable untuk menampung data provinsi
	var provinsi []models.Provinsi

	// Parsing json yang diterima dari 3rd party api
	if err := json.Unmarshal(body, &provinsi); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to unmarshal json",
			"errors":  []string{"Failed to unmarshal json"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   provinsi,
		"status": true,
	})

}

func GetKotaByProvinsiId(c *fiber.Ctx) error {
	provinsiID := c.Params("provinsiID")

	// Ambil data kota berdasarkan provinsiID dari 3rd party api
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/" + provinsiID + ".json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mendapatkan kota berdasarkan provinsi",
			"errors":  []string{"Gagal mendapatkan kota berdasarkan provinsi"},
			"data":    nil,
		})
	}

	defer resp.Body.Close()

	// Membaca body response dari 3rd party api
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membaca body response kota",
		})
	}

	// Deklarasi variable untuk menampung data kota
	var kota []models.Kota

	// Parsing json yang diterima dari 3rd party api
	if err := json.Unmarshal(body, &kota); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to unmarshal json",
			"errors":  []string{"Failed to unmarshal json"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   kota,
		"status": true,
	})
}

func GetKecamatanByKotaID(c *fiber.Ctx) error {
	kotaID := c.Params("kotaID")

	// Ambil kecamatan berdasarkan kotaID dari 3rd party api
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/districts/" + kotaID + ".json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mendapatkan kecamatan berdasarkan kota",
			"errors":  []string{"Gagal mendapatkan kecamatan berdasarkan kota"},
			"data":    nil,
		})
	}

	defer resp.Body.Close()

	// Membaca body response dari 3rd party api
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membaca body response kecamatan",
		})
	}

	// Deklarasi variable untuk menampung data kecamatan
	var kecamatan []models.Kecamatan

	// Parsing json yang diterima dari 3rd party api
	if err := json.Unmarshal(body, &kecamatan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to unmarshal json",
			"errors":  []string{"Failed to unmarshal json"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   kecamatan,
		"status": true,
	})

}

func GetKelurahanByKecamatanID(c *fiber.Ctx) error {
	kecamatanID := c.Params("kecamatanID")

	// Ambil kelurahan berdasarkan kecamatanID dari 3rd party api
	resp, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/villages/" + kecamatanID + ".json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal mendapatkan kelurahan berdasarkan kecamatan",
			"errors":  []string{"Gagal mendapatkan kelurahan berdasarkan kecamatan"},
			"data":    nil,
		})
	}

	defer resp.Body.Close()

	// Membaca body response dari 3rd party api
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Gagal membaca body response kelurahan",
		})
	}

	// Deklarasi variable untuk menampung data kelurahan
	var kelurahan []models.Kelurahan

	// Parsing json yang diterima dari 3rd party api
	if err := json.Unmarshal(body, &kelurahan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to unmarshal json",
			"errors":  []string{"Failed to unmarshal json"},
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   kelurahan,
		"status": true,
	})
}
