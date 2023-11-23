package helper

import "github.com/go-playground/validator/v10"

// Response adalah struktur data yang digunakan untuk mengembalikan respons API standar.
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta adalah struktur data yang menyimpan informasi meta-data untuk respons API.
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// ApiResponse adalah fungsi yang menghasilkan instance Response berdasarkan parameter yang diberikan.
func ApiResponse(message string, code int, status string, data interface{}) Response {
	// Membuat instance Meta berdasarkan parameter yang diberikan
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	// Membuat instance Response dengan Meta dan Data yang telah ditentukan
	response := Response{
		Meta: meta,
		Data: data,
	}

	// Mengembalikan instance Response
	return response
}

// FormatValidationError adalah fungsi yang mengonversi error validasi ke dalam bentuk slice string.
func FormatValidationError(err error) []string {
	var errors []string

	// Loop melalui setiap error validasi dan menambahkannya ke dalam slice string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	// Mengembalikan slice string yang berisi error validasi
	return errors
}
