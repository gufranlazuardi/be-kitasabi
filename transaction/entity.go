package transaction

import (
	"time"
	"tiny-donate/campaign"
	"tiny-donate/user"
)

type Transaction struct {
	ID 			int
	CampaignID	int
	UserID		int
	Amount		int
	Status 		string
	Code		string
	PaymentURL	string
	Campaign	campaign.Campaign
	User		user.User
	CreatedAt	time.Time
	UpdatedAt	time.Time
}