package campaign

import "strings"

type CampaignFormatter struct {
	ID 					int 	`json:"id"`
	UserId				int 	`json:"user_id"`
	Name				string 	`json:"name"`
	ShortDescription 	string 	`json:"short_description"`
	ImageUrl			string 	`json:"image_url"`
	GoalAmount			int 	`json:"goal_amount"`
	CurrentAmount		int 	`json:"current_amount"`
	Slug				string	`json:"slug"`
}

type CampaignDetailFormatter struct {
	ID 					int 						`json:"id"`
	Name 				string 						`json:"name"`
	ShortDescription 	string 						`json:"short_description"`
	ImageURL			string 						`json:"image_url"`
	GoalAmount 			int 						`json:"goal_amount"`
	CurrentAmmount		int							`json:"current_amount"`
	UserID				int 						`json:"user_id"`
	Slug				string 						`jsonn:"slug"`
	Perks				[]string					`json:"perks"`
	User 				CampaignUserFormatter 		`json:"user"`
	Images				[]CampaignImageFormatter	`json:"images"`
}

type CampaignUserFormatter struct {
	Name		string `json:"name"`
	ImageURL	string
}

type CampaignImageFormatter struct {
	ImageURL	string 	`json:"image_url"`
	IsPrimary	bool	`json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	// kalo campaigns kosong, return array kosong []
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter{
	var campaignDetailFormatter = CampaignDetailFormatter{}

	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.UserID = campaign.UserId
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmmount = campaign.CurrentAmount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk)) 
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User
	var campaignUserFormatter = CampaignUserFormatter{}

	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages{
		campainImageFormatter := CampaignImageFormatter{}
		campainImageFormatter.ImageURL = image.FileName

		NewIsPrimary := false

		if image.IsPrimary == 1 {
			NewIsPrimary = true
		}

		campainImageFormatter.IsPrimary = NewIsPrimary

		images = append(images, campainImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}


