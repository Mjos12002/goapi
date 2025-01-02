package sale

import "gorm.io/gorm"

type SaleModel struct {
	Product  int
	Quantity int
	UnitCost int
	gorm.Model
}
