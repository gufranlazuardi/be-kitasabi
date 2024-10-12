package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"tiny-donate/auth"
	"tiny-donate/handler"
	"tiny-donate/helper"
	"tiny-donate/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=macbook password=1234 dbname=tiny-donate port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewReposistory(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	
	testToken, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxM30.7ii5SXCsL624sd-W6AZX2F_1444oHKKZt47B_bYQ8QU")
	if err != nil {
		fmt.Println("======================")
		fmt.Println("Token error")
		fmt.Println("======================")
	}

	if testToken.Valid {
		fmt.Println("======================")
		fmt.Println("Token valid")
		fmt.Println("======================")
	}

	fmt.Println(authService.GenerateToken(1001))

	userService.SaveAvatar(5, "images/1-profile.png")
	
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions",userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}	

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	
		// Bearer tokentokentoken
		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
	
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}



// ambil nilai header Authorization: Bearer token
// dari header Authorization, ambil nilai tokennya
// validasi tokennya
// jika valid, ambil user_id
// ambil user dari db berdasarkan user_id lewat service
// set context isinya user


// controller
// handler
// service
// repository
// db