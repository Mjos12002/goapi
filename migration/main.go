package main

import (
	"advanced/model/customertype"
	"advanced/model/product"
	"advanced/model/producttype"
	"advanced/utils/database"
)

func main() {
	db := database.DbConnection()
	db.AutoMigrate(&customertype.CustomerType{}, producttype.ProductType{}, product.Product{})
}
