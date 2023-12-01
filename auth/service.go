package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// Service mendefinisikan kontrak untuk layanan otentikasi.
type Service interface {
	// GenerateToken menghasilkan JWT untuk userID yang diberikan.
	GenerateToken(userID int) (string, error)

	// ValidateToken memvalidasi JWT yang diberikan.
	ValidateToken(token string) (*jwt.Token, error)
}

// jwtService mengimplementasikan antarmuka Service untuk otentikasi berbasis JWT.
type jwtService struct {
}

// SECRET_KEY menyimpan kunci rahasia yang digunakan untuk penandatanganan JWT.
var SECRET_KEY = []byte("cob4_s3cr3tk3y")

// NewService membuat instance baru dari jwtService.
func NewService() *jwtService {
	return &jwtService{}
}

// GenerateToken menghasilkan JWT untuk userID yang diberikan.
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// Buat klaim JWT dengan user_id sebagai pasangan key-value.
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// Buat JWT baru dengan metode penandatanganan HS256 dan klaim tersebut.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Tandatangani token dengan SECRET_KEY.
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

// ValidateToken memvalidasi JWT yang diberikan.
func (s *jwtService) ValidateToken(encodedtoken string) (*jwt.Token, error) {
	// Parse token terenkripsi dengan fungsi validasi kustom.
	token, err := jwt.Parse(encodedtoken, func(token *jwt.Token) (interface{}, error) {
		// Periksa apakah metode penandatanganan adalah HMAC (HS256).
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token tidak valid")
		}
		// Kembalikan SECRET_KEY untuk validasi.
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
