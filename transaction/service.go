package transaction

import (
	"errors"
	"tiny-donate/campaign"
	"tiny-donate/payment"
)

type service struct {
	repository 			Repository
	campaignRepository 	campaign.Repository
	paymentService		payment.Service
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
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

func(s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error)  {
	transaction := Transaction{}

	transaction.Amount = input.Amount
	transaction.CampaignId = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	
	newTransaction, err := s.repository.Save(transaction)

	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)

	if err != nil {
		return paymentURL, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}