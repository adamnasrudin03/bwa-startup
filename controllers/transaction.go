package controllers

import "bwa-startup/transaction"

type transactionController struct {
	service transaction.Service
}

func NewTransactionController(service transaction.Service) *transactionController {
	return &transactionController{service}
}