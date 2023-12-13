package user

import (
	"time"

	"gorm.io/gorm"
)

// Repository adalah interface untuk operasi database pengguna.
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

// repository adalah implementasi Repository.
type repository struct {
	db *gorm.DB
}

// NewRepository membuat instance Repository dengan koneksi database yang diberikan.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Save menyimpan pengguna ke database.
func (r *repository) Save(user User) (User, error) {
	now := time.Now()
	user.CreateAt = now
	user.UpdateAt = now

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// FindByEmail mencari pengguna berdasarkan alamat email.
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

// FindByID mencari pengguna berdasarkan ID.
func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

// Update memperbarui informasi pengguna di database.
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, nil
	}

	return user, nil
}
