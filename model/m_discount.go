package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Discount struct {
	Description string
	Pct         int `gorm:"not null" binding:"required"`
	gorm.Model
	Customer []*Customer `gorm:"many2many:customer_discounts"`
}

func (Discount) TableName() string {
	return "m_discount"
}

func (d *Discount) ToString() string { // cetak seperti json
	disc, err := json.MarshalIndent(d, "", "")
	if err != nil {
		return ""
	}
	return string(disc)
}
