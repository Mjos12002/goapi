package producttype

import "gorm.io/gorm"

type ProductType struct {
	Name string
	gorm.Model
}
