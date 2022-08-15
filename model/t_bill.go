package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	CreatedAt   time.Time `gorm:"DEFAULT:current_timestamp"`
	CustomerID  int
	TableID     int
	TransTypeID int
	BillDetail  []BillDetail `gorm:"not null"`
	gorm.Model
}

func (Bill) TableName() string {
	return "t_bill"
}
func (b *Bill) ToString() string { // cetak seperti json
	bill, err := json.MarshalIndent(b, "", "")
	if err != nil {
		return ""
	}
	return string(bill)
}
