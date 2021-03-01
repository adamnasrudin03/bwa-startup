package transaction

import "bwa-startup/users"

type GetTransactionByCampaignIdInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.User
}