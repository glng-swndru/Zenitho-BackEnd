package user

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

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
