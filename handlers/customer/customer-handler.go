package customer

import (
	"advanced/model"

	"github.com/gin-gonic/gin"
)

func CustomerHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data": model.Customer{Name: "Joseph", Category: "Simple"},
	})
}
