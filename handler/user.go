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
// Metode ini mengambil input dari request, melakukan pemrosesan menggunakan service,
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

func (h *usersHandler) Login(c *gin.Context) {
	// User memasukkan input (email & password)
	// Input ditangkap handler
	// Mapping dari input user ke input struct
	// Input struct passing service
	// Di service mencari dengan bantuan repository user dengan email x
	// Mencocokkan password
}
