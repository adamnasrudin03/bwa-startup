package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetByCampaignId(campaignID int) ( []Transaction, error )

}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignID int) ( []Transaction, error ) {
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}