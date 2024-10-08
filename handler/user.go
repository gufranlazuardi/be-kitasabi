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

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors":errors}	

		response := helper.APIResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// tampung user yang login
	loggedInUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	formatter := user.FormatUser(loggedInUser, "token")
	
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
	
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	// ==============================================================================================================
	// input email dari user
	// input email di mapping ke struct input
	// struct input di passing di service
	// service akan memanggil repository untuk menentukan apakah email sudah ada apa belum
	// repository akan melakukan query ke database
	// ==============================================================================================================
	

	// tangkap inputan dari user
	var input user.EmailUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// kalo udah panggil service nya
	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 
	}

	data := gin.H{
		"is_available": IsEmailAvailable, 
	}

	var metaMessage string

	if IsEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email have been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}