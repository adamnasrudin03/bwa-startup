package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetByCampaignId(campaignID int) ( []Transaction, error )
	GetByUserId(userID int) ([]Transaction, error)

}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignID int) ( []Transaction, error ) {
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserId(userID int) ([]Transaction, error) {
	var transactions []Transaction
	//preload [load data realasi ke campaign lalu ke campaign images, yg di load
	// di campaign images is_primary == 1]
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary == 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}