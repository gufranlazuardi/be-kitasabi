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

	userByEmail, err := userRepository.FindByEmail("balmondddd@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if (userByEmail.ID == 0){
		fmt.Println("User not found")
	} else {
		fmt.Println(userByEmail.Name)
	}

	
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