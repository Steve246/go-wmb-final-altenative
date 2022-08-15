package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Menu struct {
	MenuName  string `gorm:"size:100;not null" binding:"required"`
	MenuPrice MenuPrice
	gorm.Model
}

func (Menu) TableName() string {
	return "m_menu"
}

func (m *Menu) ToString() string { // cetak seperti json
	menu, err := json.MarshalIndent(m, "", "")
	if err != nil {
		return ""
	}
	return string(menu)
}
