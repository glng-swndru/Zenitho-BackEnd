package helper

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
