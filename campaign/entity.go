package campaign

import (
	"time"
	"tiny-donate/user"
)

type Campaign struct {
	ID 					int
	UserId				int
	Name				string
	ShortDescription 	string
	LongDescription 	string
	Perks				string
	DonorCount			int
	GoalAmount			int
	CurrentAmount		int
	Slug				string
	CreatedAt 			time.Time	
	UpdatedAt 			time.Time
	CampaignImages		[]CampaignImage
	User				user.User
}

type CampaignImage struct {
	ID				int
	CampaignId		int
	FileName		string
	IsPrimary		int
	CreatedAt		time.Time
	UpdatedAt		time.Time
}