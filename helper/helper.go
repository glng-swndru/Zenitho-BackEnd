package helper

import "github.com/go-playground/validator/v10"

// Response adalah struktur data yang digunakan untuk mengembalikan respons API standar.
type Response struct {
	Meta Meta        `json:"meta"` // Bagian meta dari respons, berisi info umum seperti pesan, kode, dan status.
	Data interface{} `json:"data"` // Bagian data dari respons, bisa berisi apa saja.
}

// Meta adalah struktur data yang menyimpan informasi meta-data untuk respons API.
type Meta struct {
	Message string `json:"message"` // Pesan dalam respons, seperti "Berhasil memuat data".
	Code    int    `json:"code"`    // Kode status HTTP, seperti 200, 400, 404, dll.
	Status  string `json:"status"`  // Status operasi, biasanya "success" atau "error".
}

// ApiResponse adalah fungsi yang menghasilkan instance Response berdasarkan parameter yang diberikan.
func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message, // Set pesan.
		Code:    code,    // Set kode status HTTP.
		Status:  status,  // Set status operasi.
	}

	response := Response{
		Meta: meta, // Set bagian meta dari respons.
		Data: data, // Set bagian data dari respons.
	}

	return response // Kembalikan respons yang sudah dibuat.
}

// FormatValidationError adalah fungsi yang mengonversi error validasi ke dalam bentuk slice string.
func FormatValidationError(err error) []string {
	var errors []string // Siapin slice untuk tampung pesan error.

	// Loop melalui setiap error validasi.
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error()) // Tambahkan pesan error ke slice.
	}

	return errors // Kembalikan slice berisi pesan error.
}
