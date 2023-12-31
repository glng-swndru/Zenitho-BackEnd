package user

// RegisterUserInput adalah struktur data yang digunakan sebagai input saat mendaftarkan pengguna baru.
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// LoginInput adalah struktur data yang digunakan sebagai input saat pengguna melakukan login.
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CheckEmailInput adalah struktur data yang digunakan sebagai input saat memeriksa ketersediaan alamat email.
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
