package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName  string      `gorm:"size:100;not null" binding:"required"`
	MobilePhoneNo string      `gorm:"unique" binding:"required"`
	IsMember      bool        `gorm:"default:false"`
	Discount      []*Discount `gorm:"many2many:customer_discounts"`
	Bill          []Bill
}

func (Customer) TableName() string {
	return "m_customer"
}
func (c *Customer) ToString() string { // cetak seperti json
	cust, err := json.MarshalIndent(c, "", "")
	if err != nil {
		return ""
	}
	return string(cust)
}
