package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service adalah antarmuka yang mendefinisikan operasi-operasi yang dapat dilakukan terhadap entitas pengguna.
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

// service adalah implementasi dari antarmuka Service.
type service struct {
	repository Repository
}

// NewService digunakan untuk membuat instance baru dari Service dengan repository yang diberikan.
func NewService(repository Repository) *service {
	return &service{repository}
}

// RegisterUser merupakan metode untuk mendaftarkan pengguna baru.
// Metode ini mengambil input dari RegisterUserInput, menghasilkan User baru, dan menyimpannya ke repository.
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// Membuat instance User baru
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	// Menghasilkan hash dari kata sandi menggunakan bcrypt
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	// Menyimpan pengguna baru ke repository
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	// Mengembalikan pengguna baru setelah berhasil mendaftar
	return newUser, nil
}

// Login merupakan metode untuk melakukan proses login pengguna.
// Metode ini mengambil input dari LoginInput, mencari pengguna berdasarkan alamat email,
// dan memeriksa kesesuaian password menggunakan bcrypt.
func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	// Mencari pengguna berdasarkan alamat email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return User{}, err
	}

	// Kembalikan error jika tidak ada pengguna dengan alamat email yang diberikan
	if user.ID == 0 {
		return User{}, errors.New("no user found on that email")
	}

	// Memeriksa kesesuaian password menggunakan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		// Kembalikan error jika password tidak cocok
		return User{}, err
	}

	// Kembalikan user jika login berhasil
	return user, nil
}

// IsEmailAvailable merupakan metode untuk memeriksa ketersediaan alamat email.
// Metode ini mengambil input dari CheckEmailInput, mencari pengguna berdasarkan alamat email,
// dan mengembalikan informasi ketersediaan alamat email.
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	// Mencari pengguna berdasarkan alamat email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// Kembalikan true jika tidak ada pengguna dengan alamat email yang diberikan
	if user.ID == 0 {
		return true, nil
	}

	return true, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}