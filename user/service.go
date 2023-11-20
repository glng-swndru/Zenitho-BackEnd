package user

import "golang.org/x/crypto/bcrypt"

// Service merupakan antarmuka yang mendefinisikan operasi-operasi yang dapat dilakukan terhadap entitas pengguna (user).
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
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
