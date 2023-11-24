package handler

import (
	"campaignku/helper"
	"campaignku/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// usersHandler adalah struct yang menyediakan metode-handler untuk entitas pengguna (user).
type usersHandler struct {
	userService user.Service
}

// NewUserHandler digunakan untuk membuat instance baru dari usersHandler dengan service yang diberikan.
func NewUserHandler(userService user.Service) *usersHandler {
	return &usersHandler{userService}
}

// RegisterUser adalah metode-handler untuk mendaftarkan pengguna baru.
// Metode ini mengambil input dari request, melakukan validasi, pemrosesan menggunakan service,
// dan mengembalikan respons API sesuai hasil pemrosesan.
func (h *usersHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	// Tangkap dan validasi input dari user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Mendaftarkan pengguna menggunakan service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "success", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Memformat data pengguna untuk respons API
	formatter := user.FormatUser(newUser, "tokentokentoken")

	// Membuat respons API
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	// Mengembalikan respons API
	c.JSON(http.StatusOK, response)
}

// Login adalah metode-handler untuk proses login pengguna.
// Metode ini mengambil input dari request, melakukan validasi, pemrosesan menggunakan service,
// dan mengembalikan respons API sesuai hasil pemrosesan.
func (h *usersHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Melakukan login menggunakan service
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Memformat data pengguna untuk respons API
	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	// Membuat respons API
	response := helper.ApiResponse("Successfully logged in", http.StatusOK, "success", formatter)

	// Mengembalikan respons API
	c.JSON(http.StatusOK, response)
}

// CheckEmailAvailability adalah metode-handler untuk memeriksa ketersediaan alamat email.
// Metode ini mengambil input dari request, melakukan validasi, pemrosesan menggunakan service,
// dan mengembalikan respons API sesuai hasil pemrosesan.
func (h *usersHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Memeriksa ketersediaan alamat email menggunakan service
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Menyiapkan data untuk respons API
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	// Menyiapkan pesan meta berdasarkan ketersediaan alamat email
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	// Membuat respons API
	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
