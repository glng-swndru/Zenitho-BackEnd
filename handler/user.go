// Package handler menyediakan fungsi-fungsi penanganan permintaan HTTP untuk entitas pengguna dalam aplikasi Campaignku.
package handler

import (
	"campaignku/auth"
	"campaignku/helper"
	"campaignku/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// usersHandler adalah tipe data yang menyediakan fungsi-fungsi penanganan permintaan terkait pengguna.
type usersHandler struct {
	userService user.Service // Layanan pengguna.
	authService auth.Service // Layanan otentikasi.
}

// NewUserHandler membuat objek usersHandler baru dengan layanan pengguna dan otentikasi yang diperlukan.
func NewUserHandler(userService user.Service, authService auth.Service) *usersHandler {
	return &usersHandler{userService, authService} // Return instance usersHandler.
}

// RegisterUser menangani permintaan pendaftaran pengguna baru.
func (h *usersHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput // Siapin variabel buat input dari user.

	// Ambil dan validasi input dari pengguna.
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// Handle error validasi.
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Gagal mendaftarkan akun", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Daftarkan pengguna menggunakan layanan.
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		// Handle error saat registrasi.
		response := helper.ApiResponse("Gagal mendaftarkan akun", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Generate token JWT setelah registrasi sukses.
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Gagal mendaftarkan akun", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Format data pengguna dan token.
	formatter := user.FormatUser(newUser, token)
	response := helper.ApiResponse("Akun berhasil didaftarkan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// Login menangani permintaan login pengguna.
func (h *usersHandler) Login(c *gin.Context) {
	var input user.LoginInput // Siapin variabel buat input dari user.

	// Ambil dan validasi input untuk login.
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// Handle error validasi.
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Lakukan login menggunakan layanan.
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		// Handle error saat login.
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Generate token JWT setelah login sukses.
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ApiResponse("Login gagal", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Format data pengguna dan token.
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.ApiResponse("Berhasil login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// CheckEmailAvailability menangani permintaan pengecekan ketersediaan alamat email.
func (h *usersHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput // Siapin variabel buat input.

	// Ambil dan validasi input untuk cek email.
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// Handle error validasi.
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Pengecekan email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Periksa ketersediaan alamat email menggunakan layanan.
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		// Handle error saat cek ketersediaan email.
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Pengecekan email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Siapin data respons.
	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email telah terdaftar"
	if isEmailAvailable {
		metaMessage = "Email tersedia"
	}

	// Kirim respons.
	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// UploadAvatar menangani permintaan pengunggahan avatar pengguna.
func (h *usersHandler) UploadAvatar(c *gin.Context) {
	// Ambil file avatar dari form-data.
	file, err := c.FormFile("avatar")
	if err != nil {
		// Handle error saat mengambil file.
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengunggah gambar avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Dapatkan info pengguna saat ini.
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// Tentukan path penyimpanan file.
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	// Simpan file yang diunggah.
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		// Handle error saat menyimpan file.
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengunggah gambar avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Update path avatar di database.
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		// Handle error saat update database.
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengunggah gambar avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Kirim respons sukses.
	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Avatar berhasil diunggah", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
