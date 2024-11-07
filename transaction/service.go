package transaction

import (
	"errors"
	"strconv"
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
	ProcessPayment(input TransactionNotificationInput) (error)
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
	transaction.CampaignID = input.CampaignID
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
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}


// handle notification from midtrans
func(s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID) 

	transaction, err := s.repository.GetByID(transaction_id)

	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expired" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return nil          
	}

	// update data campaign
	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.DonorCount = campaign.DonorCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount


		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	
	return nil
}