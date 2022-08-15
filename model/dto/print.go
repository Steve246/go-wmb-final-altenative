package dto

import (
	"encoding/json"
	"time"
)

type CetakBill struct {
	TansDate  time.Time
	Customers string
	Table     string
	Transtype string
	// Detail    []CetakBillDetail
}

func (tr *CetakBill) ToString() string { // cetak seperti json
	trType, err := json.MarshalIndent(tr, "", "")
	if err != nil {
		return ""
	}
	return string(trType)
}
