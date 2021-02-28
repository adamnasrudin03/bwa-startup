package controllers

import (
	"bwa-startup/campaign"
	"bwa-startup/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type campaignController struct {
	service campaign.Service
}

func NewCampaignController(service campaign.Service) *campaignController {
	return &campaignController{service}
}

func (h *campaignController) GetCampaigns(c *gin.Context) {
	//mengubah string parameter user_id ke int, dgn return nilai, error
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helpers.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}
