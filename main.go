package main

import (
	"fmt"
	"log"
	"tiny-donate/auth"
	"tiny-donate/handler"
	"tiny-donate/user"

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

	userService.SaveAvatar(1, "images/1-profile.png")
	
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions",userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}	


// controller
// handler
// service
// repository
// db