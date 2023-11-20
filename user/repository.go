package user

import (
	"time"

	"gorm.io/gorm"
)

// Repository adalah antarmuka yang mendefinisikan operasi-operasi penyimpanan (database) terhadap entitas pengguna (user).
type Repository interface {
	Save(user User) (User, error)
}

// repository adalah implementasi dari antarmuka Repository.
type repository struct {
	db *gorm.DB
}

// NewRepository digunakan untuk membuat instance baru dari Repository dengan koneksi database yang diberikan.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save digunakan untuk menyimpan pengguna ke dalam database.
// Metode ini mengambil instance User, menetapkan waktu pembuatan dan pembaruan, dan menyimpannya ke dalam database.
func (r *repository) Save(user User) (User, error) {
	// Mendapatkan waktu sekarang
	now := time.Now()

	// Menetapkan waktu pembuatan dan pembaruan pengguna
	user.CreateAt = now
	user.UpdateAt = now

	// Menyimpan pengguna ke dalam database menggunakan GORM
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	// Mengembalikan pengguna setelah berhasil disimpan
	return user, nil
}
