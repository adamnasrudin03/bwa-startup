package controllers

import (
	"bwa-startup/helpers"
	"bwa-startup/transaction"
	"bwa-startup/users"
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
	
	currentUser := c.MustGet("currentUser").(users.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helpers.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of Campaigns transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) GetUserTransaction (c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userId := currentUser.ID

	transactions, err := h.service.GetTransactionByUserId(userId)
	if err != nil {
		response := helpers.APIResponse("Failed to get users transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of users transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)

}
