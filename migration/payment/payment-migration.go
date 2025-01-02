package main

import (
	"advanced/model/payment"
	"advanced/utils/database"
)

func main() {
	dt := database.DbConnection()
	dt.AutoMigrate(&payment.Payment{})
}
