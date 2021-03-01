package transaction

import (
	"bwa-startup/campaign"
	"bwa-startup/users"
	"time"
)


type Transaction struct {
	ID 			int
	CampaignID 	int
	UserID		int
	Amount 		int
	Status 		string
	Code 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	User        users.User
	Campaign    campaign.Campaign
}