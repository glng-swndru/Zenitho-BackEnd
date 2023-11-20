package user

// RegisterUserInput adalah struktur data yang digunakan sebagai input saat mendaftarkan pengguna baru.
type RegisterUserInput struct {
	Name       string // Name adalah nama lengkap dari pengguna yang akan didaftarkan.
	Occupation string // Occupation adalah pekerjaan atau profesi dari pengguna yang akan didaftarkan.
	Email      string // Email adalah alamat email unik dari pengguna yang akan didaftarkan.
	Password   string // Password adalah kata sandi yang akan digunakan oleh pengguna yang akan didaftarkan.
}
