package main

import (
	"log"
	"tiny-donate/user"

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
	
	userInput := user.RegisterUserInput{}
	userInput.Name = "gifania dari service"
	userInput.Email = "bonjoganteng@gmail.com"
	userInput.Occupation = "anak presiden"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}	


// controller
// handler
// service
// repository
// db