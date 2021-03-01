package transaction

import "bwa-startup/users"

type GetTransactionByCampaignIdInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.User
}

type CreateTransactionInput struct {
	Amount   	int `uri:"amount" binding:"required"`
	CampaignID  int `uri:"campaign_id" binding:"required"`
	User 		users.User
}