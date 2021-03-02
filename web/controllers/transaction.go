package controllers

import (
	"bwa-startup/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)


type transactionController struct {
	transactionService transaction.Service
}

func NewTransactionController(transactionService transaction.Service) *transactionController {
	return &transactionController{transactionService}
}

func (h *transactionController) Index(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransaction()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}