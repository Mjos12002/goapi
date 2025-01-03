package main

import (
	"advanced/handlers"
	"advanced/handlers/customer"
	userhander "advanced/handlers/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	router := gin.Default()

	router.GET("/api", handlers.WelcomeHandler)

	router.GET("/customer", customer.CustomerHandler)

	router.GET("/user", userhander.UserHandler)

	router.POST("/api/customer", customer.CreateCustomerHandler)

	err := router.Run()
	if err != nil {
		panic(err)
	}

}
