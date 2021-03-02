package transaction

import (
	"bwa-startup/campaign"
	"bwa-startup/users"
	"time"

	"github.com/leekchan/accounting"
)


type Transaction struct {
	ID 			int
	CampaignID 	int
	UserID		int
	Amount 		int
	Status 		string
	Code 		string
	PaymentURL	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	User        users.User
	Campaign    campaign.Campaign
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
