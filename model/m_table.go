package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Table struct {
	TableDescription string `gorm:"size:50;not null;unique" binding:"required"`
	IsAvailable      bool   `gorm:"default:true"`
	Bill             []Bill
	gorm.Model
}

func (Table) TableName() string {
	return "m_table"
}

func (t *Table) ToString() string { // cetak seperti json
	table, err := json.MarshalIndent(t, "", "")
	if err != nil {
		return ""
	}
	return string(table)
}
