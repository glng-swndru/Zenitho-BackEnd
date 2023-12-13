package user

// UserFormatter adalah struktur data yang digunakan untuk memformat data pengguna (user) sebelum dikirim sebagai respons API.
type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

// FormatUser adalah fungsi yang menghasilkan instance UserFormatter berdasarkan instance User dan token yang diberikan.
func FormatUser(user User, token string) UserFormatter {
	// Membuat instance UserFormatter dengan menggunakan data dari instance User dan token yang diberikan.
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	// Mengembalikan instance UserFormatter yang telah diformat.
	return formatter // Kembalikan data pengguna yang telah diformat.
}
