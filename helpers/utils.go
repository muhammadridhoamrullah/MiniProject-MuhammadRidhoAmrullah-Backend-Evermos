package helpers

import (
	"regexp"
	"strings"
)

func GenerateSlug(namaProduk string) string {
	// Mengubah nama produk menjadi huruf kecil
	slug := strings.ToLower(namaProduk)

	// Mengganti spasi dengan tanda hubung
	slug = strings.ReplaceAll(slug, " ", "-")

	// Menghapus karakter yang tidak alphanumeric (termasuk tanda baca dan simbol)
	re := regexp.MustCompile(`[^a-z0-9-]`)
	slug = re.ReplaceAllString(slug, "")

	return slug
}

//
