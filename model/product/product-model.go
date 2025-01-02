package product

import "gorm.io/gorm"

type Product struct {
	Name string
	gorm.Model
}
