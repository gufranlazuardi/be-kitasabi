package transaction

import (
	"time"
	"tiny-donate/campaign"
	"tiny-donate/user"
)

type Transaction struct {
	ID 			int
	CampaignId	int
	UserID		int
	Amount		int
	Status 		string
	Code		string
	Campaign	campaign.Campaign
	User		user.User
	CreatedAt	time.Time
	UpdatedAt	time.Time
}