package handler

import (
	"net/http"
	"tiny-donate/helper"
	"tiny-donate/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}	

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita pasing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		// mapping error datanya di response kalo gagal (contoh: password required, email format, dll) manggil di helper
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors":errors,
		}	

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	// token, err

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}		

func(h *userHandler) LoginUser(c *gin.Context) {
	// user memasukkan input (email, password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct nya passing ke service
	// di service mencari dengan bantuan repository user dengan email user
	// jika ketemu, validasi password
}