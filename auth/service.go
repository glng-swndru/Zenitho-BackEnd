package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// Service mendefinisikan 'kontrak' kerja untuk layanan otentikasi.
type Service interface {
	GenerateToken(userID int) (string, error)       // Fungsi buat bikin token JWT dari userID.
	ValidateToken(token string) (*jwt.Token, error) // Fungsi buat cek token JWT itu valid apa enggak.
}

// jwtService, implementasi dari Service, spesial buat JWT.
type jwtService struct {
	// Nggak ada field khusus di sini.
}

// SECRET_KEY, kunci rahasia buat tanda tangan digital di JWT.
var SECRET_KEY = []byte("cob4_s3cr3tk3y")

// NewService buat instance baru jwtService.
func NewService() *jwtService {
	return &jwtService{} // Kembalikan struct jwtService yang baru dibuat.
}

// GenerateToken bikin token JWT dari userID.
func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}  // Bikin 'klaim' buat token.
	claim["user_id"] = userID // Masukin userID ke dalam klaim.

	// Bikin token baru dengan metode HS256, masukin klaim tadi.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY) // Tanda tangani token pake SECRET_KEY.
	if err != nil {
		return signedToken, err // Kalo ada error, balikin errornya.
	}
	return signedToken, nil // Kalo sukses, balikin token yang udah ditanda tangani.
}

// ValidateToken cek apakah token JWT yang diberikan itu valid.
func (s *jwtService) ValidateToken(encodedtoken string) (*jwt.Token, error) {
	// Parse token, pake fungsi kustom buat validasi.
	token, err := jwt.Parse(encodedtoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // Cek metodenya HMAC (HS256) apa enggak.

		if !ok {
			return nil, errors.New("token tidak valid") // Kalo bukan, bilang tokennya nggak valid.
		}
		return []byte(SECRET_KEY), nil // Kalo iya, balikin SECRET_KEY buat validasi.
	})

	if err != nil {
		return token, err // Kalo parsing error, balikin errornya.
	}
	return token, nil // Kalo sukses, balikin token yang udah divalidasi.
}
