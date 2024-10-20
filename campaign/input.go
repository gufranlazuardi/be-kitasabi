package campaign

import "tiny-donate/user"

type GetCampaignDetailInput struct {
	ID 	int `uri:"id" binding:"required"` 
}

type CreateCampaignInput struct {
	Name 				string 		`json:"name" binding:"required"`
	ShortDescription 	string 		`json:"short_description" binding:"required"`
	LongDescription 	string 		`json:"long_description" binding:"required"`
	GoalAmount			int 		`json:"goal_amount" binding:"required"`
	Perks				string 		`json:"perks" binding:"required"`
	User				user.User
}

type UpdateCampaignInput struct {
	Name 				string 		`json:"name" binding:"required"`
	ShortDescription 	string 		`json:"short_description" binding:"required"`
	LongDescription 	string 		`json:"long_description" binding:"required"`
	GoalAmount			int 		`json:"goal_amount" binding:"required"`
	Perks				string 		`json:"perks" binding:"required"`
	User				user.User
}