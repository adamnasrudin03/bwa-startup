package controllers

import (
	"bwa-startup/helpers"
	"bwa-startup/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)


type transactionController struct {
	service transaction.Service
}

func NewTransactionController(service transaction.Service) *transactionController {
	return &transactionController{service}
}

func (h *transactionController) GetCampaignTransaction (c *gin.Context) {
	var input transaction.GetTransactionByCampaignIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helpers.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Campaign's transaction", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}
