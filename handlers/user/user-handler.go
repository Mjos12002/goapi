package user

import (
	"advanced/model/product"

	"github.com/gin-gonic/gin"
)

func UserHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": product.Product{Name: "Ibirayi"},
	})
}
