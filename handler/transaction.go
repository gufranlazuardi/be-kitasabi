package handler

import (
	"net/http"
	"tiny-donate/helper"
	"tiny-donate/transaction"
	"tiny-donate/user"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter mapping input struct (pake should bind uri)
// panggil service, input struct sebagai parameternya
// service punya campaign_id

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func(h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	// tangkap inputnya dulu
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get campaign's transactions success", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func(h *transactionHandler) GetUserTransactions(c *gin.Context) {
	// ambil siapa user yang melakukan request

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get users's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Users's transactions success", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func(h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors":errors,
		}	

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newTransactions, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Create transactions success", http.StatusOK, "success", transaction.FormatTransaction(newTransactions))
	c.JSON(http.StatusOK, response)
}

func(h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProcessPayment(input)

	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
	
}


// get user transaction (siapa yang transaksi)
// handler
// ambil nilai user dari jwt
// service
// ambil data transactions (preload campaign)


// midtrans
// input dari user
// handler tangkap inputannya, mapping ke input struct
// panggil service buat transaksi, manggil sistem midtrans
// service panggil repository