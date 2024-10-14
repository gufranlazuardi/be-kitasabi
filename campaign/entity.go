package campaign

import "time"

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
}

type CampaignImages struct {
	ID				int
	CampaignId		int
	FileName		string
	IsPrimary		int
	CreatedAt		time.Time
	UpdatedAt		time.Time
}