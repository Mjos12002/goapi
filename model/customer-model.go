package model

import "gorm.io/gorm"

type Customer struct {
	Name     string
	Category string
	gorm.Model
}
