package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() gorm.DB {
	dsn := "host=localhost user=parkingapp password=123ParkingApp456 dbname=parkingapp port=5432 sslmode=disable TimeZone=Africa/Kigali"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return *db
}
