package user

import "time"

// User adalah struktur data yang merepresentasikan entitas pengguna (user).
type User struct {
	ID             int       // ID merupakan identifier unik untuk setiap pengguna.
	Name           string    // Name adalah nama lengkap dari pengguna.
	Occupation     string    // Occupation adalah pekerjaan atau profesi dari pengguna.
	Email          string    // Email adalah alamat email unik dari pengguna.
	PasswordHash   string    // PasswordHash adalah hash dari kata sandi pengguna.
	AvatarFileName string    // AvatarFileName adalah nama file avatar pengguna.
	Role           string    // Role menentukan peran atau hak akses pengguna dalam sistem.
	CreateAt       time.Time `gorm:"column:created_at"` // CreateAt adalah timestamp waktu ketika pengguna dibuat.
	UpdateAt       time.Time `gorm:"column:updated_at"` // UpdateAt adalah timestamp waktu ketika pengguna terakhir kali diperbarui.
}
