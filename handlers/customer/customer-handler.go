package customer

import (
	"advanced/model"
	"advanced/utils/database"
	"net/http"

	customerRespons "advanced/response/customer"

	"github.com/gin-gonic/gin"
)

func CustomerHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data": model.Customer{Name: "Joseph", Category: "Simple"},
	})
}

func CreateCustomerHandler(ctx *gin.Context) {
	var customer model.Customer
	if err := ctx.BindJSON(&customer); err != nil {
		return
	}
	create := database.DbConnection()
	create.Create(&customer)
	if create.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Error": create.Error,
		})
	}
	var response customerRespons.ResponseInfo
	response.Code = 200
	response.Message = "Created successfully"
	response.Status = "Success"
	response.Data = customer

	ctx.IndentedJSON(http.StatusCreated, response)
}
