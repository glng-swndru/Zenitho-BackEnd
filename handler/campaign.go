package handler

import (
	"campaignku/campaign"
	"campaignku/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Struct buat handle campaign.
type campaignHandler struct {
	service campaign.Service // Service ini yang bakal urusin logika bisnis.
}

// Fungsi buat bikin handler campaign baru.
func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service} // Balikin struct campaignHandler yang baru.
}

// Method buat dapetin data campaign.
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// Ambil user_id dari query, terus ubah jadi integer.
	userID, _ := strconv.Atoi(c.Query("user_id"))

	// Ambil data campaign dari service pake userID yang udah diambil.
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		// Kalo ada error, balikin response error.
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Kalo sukses, balikin response berisi daftar campaign.
	response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}
