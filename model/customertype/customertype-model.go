package customertype

import "gorm.io/gorm"

type CustomerType struct {
	Name string
	gorm.Model
}
