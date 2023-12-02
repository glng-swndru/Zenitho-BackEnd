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
	userService user.Service
	authService auth.Service
}

// NewUserHandler membuat objek usersHandler baru dengan menyediakan layanan pengguna dan otentikasi yang diperlukan.
func NewUserHandler(userService user.Service, authService auth.Service) *usersHandler {
	return &usersHandler{userService, authService}
}

// RegisterUser menangani permintaan pendaftaran pengguna baru.
func (h *usersHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	// Ambil dan validasi input dari pengguna
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Gagal mendaftarkan akun", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Daftarkan pengguna menggunakan layanan
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Gagal mendaftarkan akun", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Gagal mendaftarkan akun", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.ApiResponse("Akun berhasil didaftarkan", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// Login menangani permintaan login pengguna.
func (h *usersHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Lakukan login menggunakan layanan
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ApiResponse("Login gagal", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.ApiResponse("Berhasil login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// CheckEmailAvailability menangani permintaan pengecekan ketersediaan alamat email.
func (h *usersHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Pengecekan email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Periksa ketersediaan alamat email menggunakan layanan
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Pengecekan email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email telah terdaftar"
	if isEmailAvailable {
		metaMessage = "Email tersedia"
	}

	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// UploadAvatar menangani permintaan pengunggahan avatar pengguna.
func (h *usersHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengunggah gambar avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengunggah gambar avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengunggah gambar avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Avatar berhasil diunggah", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
