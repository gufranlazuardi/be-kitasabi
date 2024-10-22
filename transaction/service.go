package transaction

import (
	"errors"
	"tiny-donate/campaign"
)

type service struct {
	repository Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {

	// GET Campaign
	// Check campaign.userId != user id yang melakukan request

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignID(input.ID)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
