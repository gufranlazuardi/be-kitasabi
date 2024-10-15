package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindById(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewReposistory(db *gorm.DB) *repository {
 	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var AllCampaign []Campaign

	err := r.db.Find(&AllCampaign).Error
	if err != nil {
		return AllCampaign, err
	}

	return AllCampaign, nil
}

func (r *repository) FindById(userID int) ([]Campaign, error) {
	var AllCampaign []Campaign
	// karena punya relasi ke CampaignImages, kasih preload
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&AllCampaign).Error
	if err != nil {
		return AllCampaign, err
	}

	return AllCampaign, nil
}