package handlers

import "github.com/gin-gonic/gin"

func WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to our Country",
	})
}
