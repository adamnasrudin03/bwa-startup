package transaction

type GetTransactionByCampaignIdInput struct {
	ID int `uri:"id" binding:"required"`
}