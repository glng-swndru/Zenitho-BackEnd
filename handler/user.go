package handler

import (
	"campaignku/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usersHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *usersHandler {
	return &usersHandler{userService}
}

func (h *usersHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas di passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, user)
}
