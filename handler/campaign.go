package handler

import (
	"net/http"
	"strconv"
	"tiny-donate/campaign"
	"tiny-donate/helper"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service yang menentukan apakah repository mana yang di call
// repository :FindAll, FindUserById
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler{
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// karena yang kita tangkap user_id
	// karena user_id int dan c query return string, convert dulu

	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return 
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler)  GetCampaign(c *gin.Context){
	// handler mapping id yang di url ke struct input => service, call formatter
	// service : struct input untuk menangkap id url pake shouldbind
	// repository : get campaign by id
}

// api/v1/campaigns