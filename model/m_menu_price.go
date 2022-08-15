package model

import "gorm.io/gorm"

type MenuPrice struct {
	MenuID     int
	Price      float64 `gorm:"not null"`
	BillDetail []BillDetail
	gorm.Model
}

func (MenuPrice) TableName() string {

	return "m_menu_price"
}
