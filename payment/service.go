package payment

import (
	"os"
	"strconv"
	"tiny-donate/campaign"
	"tiny-donate/transaction"
	"tiny-donate/user"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	transactionRepository transaction.Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error) 
	ProcessPayment(input transaction.TransactionNotificationInput) (error)
}

func NewService(transactionRepository transaction.Repository, campaignRepository campaign.Repository) *service {

	return &service{transactionRepository, campaignRepository}
}

func(s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient :=  midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
    midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway {
        Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}
	
	snapTokenResp, err := snapGateway.GetToken(snapReq)

	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil

}	

// handle notification from midtrans

func(s *service) ProcessPayment(input transaction.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID) 

	transaction, err := s.transactionRepository.GetByID(transaction_id)

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

	updatedTransaction, err := s.transactionRepository.Update(transaction)
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