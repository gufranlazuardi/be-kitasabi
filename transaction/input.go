package transaction

import "tiny-donate/user"

type GetCampaignTransactionInput struct {
	ID 		int `uri:"id" binding:"required"`
	User 	user.User
}