package user

import (
	"time"

	"gorm.io/gorm"
)

// Repository adalah antarmuka yang mendefinisikan operasi-operasi penyimpanan (database) terhadap entitas pengguna (user).
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
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

// FindByEmail digunakan untuk mencari pengguna berdasarkan alamat email.
// Metode ini mengambil alamat email sebagai parameter, mencari pengguna dengan alamat email yang sesuai,
// dan mengembalikan instance User jika ditemukan.
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, nil
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, nil
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, nil
	}
	return user, nil
}
