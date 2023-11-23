package user

import "time"

// User adalah struktur data yang merepresentasikan entitas pengguna (user).
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreateAt       time.Time `gorm:"column:created_at"`
	UpdateAt       time.Time `gorm:"column:updated_at"`
}
