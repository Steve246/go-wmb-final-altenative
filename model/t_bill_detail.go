package model

import (
	"gorm.io/gorm"
)

type BillDetail struct {
	BillID      int
	MenuPriceID int
	Qty         float64 `binding:"required"`
	gorm.Model
}

func (BillDetail) TableName() string {
	return "t_bill_detail"
}
