package payment

import "gorm.io/gorm"

type Payment struct {
	Amount   int
	Customer int
	gorm.Model
}
