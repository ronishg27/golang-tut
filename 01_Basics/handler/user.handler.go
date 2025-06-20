package handler

import (
	"ronishg27/basics/database"
	"ronishg27/basics/models"
	"ronishg27/basics/utils"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request body",
			"error":   err.Error()})
		return
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"message": "failed to create user",
			"error":   result.Error.Error()})
		return
	}
	ctx.JSON(201, gin.H{
		"message": "user created successfully",
		"data":    user})
}

func LoginHandler(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.ShouldBindJSON(&loginData)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request body",
			"error":   err.Error()})
		return
	}
	loginData.Password, err = utils.HashPassword(loginData.Password)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid password",
			"error":   err.Error()})
		return
	}
	database.DB.First(loginData)

}
