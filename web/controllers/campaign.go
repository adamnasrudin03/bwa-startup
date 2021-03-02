package controllers

import (
	"bwa-startup/campaign"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignController struct {
	campaignService campaign.Service
}

func NewcampaignController(campaignService campaign.Service) *campaignController {
	return &campaignController{campaignService}
}

func (h *campaignController) Index(c *gin.Context) {
	campaigns, err := h.campaignService.GetCampaigns(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"campaigns": campaigns})
}