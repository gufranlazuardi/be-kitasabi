package main

import (
	"fmt"
	"log"
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


	input := user.LoginUserInput{
		Email: "balmond@gmail.com",
		Password: "$2a$04$s5QR16HpYgxS543oqLzbCuuGmFaUN7rqcZq9W8nyoW4KX6JcqhPEa",
	}
	user, err := userService.LoginUser(input)
	if err != nil {
		fmt.Println("========================")
		fmt.Println("There something wrong")
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)
	
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}	


// controller
// handler
// service
// repository
// db