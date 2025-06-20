package main

import (
	"log"
	"ronishg27/basics/database"
	"ronishg27/basics/handler"
	"ronishg27/basics/models"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		return
	}
	g := gin.Default()

	// setup routes
	g.POST("/user", handler.CreateUserHandler)

	err = g.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

}
