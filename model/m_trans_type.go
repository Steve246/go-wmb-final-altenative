package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type TransType struct {
	Description string `binding:"required"`
	Bill        []Bill
	gorm.Model
}

func (TransType) TableName() string {
	return "m_trans_type"
}
func (tr *TransType) ToString() string { // cetak seperti json
	trType, err := json.MarshalIndent(tr, "", "")
	if err != nil {
		return ""
	}
	return string(trType)
}
