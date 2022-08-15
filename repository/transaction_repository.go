package repository

import (
	"errors"
	"livecode-wmb-2/model"
	"livecode-wmb-2/model/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(bill *model.Bill) error
	ReadAll() ([]model.Bill, error)
	FindByIDPreload(by map[string]interface{}) (dto.CetakBill, error)
	ReadById(id int) (model.Bill, error)
	Update(bill *model.Bill, by map[string]interface{}) error
	Delete(bill *model.Bill) error
}

type transactionRepository struct {
	db *gorm.DB
}

// ReadById implements TransactionRepository
func (trx *transactionRepository) ReadById(id int) (model.Bill, error) {
	var bill model.Bill
	result := trx.db.First(&bill, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

// FindByIDPreload implements TransactionRepository
func (trx *transactionRepository) FindByIDPreload(by map[string]interface{}) (dto.CetakBill, error) {
	var transaction model.Bill
	var transactions dto.CetakBill
	result := trx.db.Model(&transaction).Select("t_bill.created_at as transDate, m_customer.customer_name as customers, m_table.table_description as table, m_trans_type.description as transtype").Joins("JOIN m_customer ON m_customer.id = t_bill.customer_id").Joins("JOIN m_table ON m_table.id = t_bill.table_id").Joins("JOIN m_trans_type ON m_trans_type.id = t_bill.trans_type_id").Joins("JOIN t_bill_detail ON t_bill_detail.id=t_bill.id").Where(by).Scan(&transactions)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transactions, nil
		} else {
			return transactions, err
		}
	}

	return transactions, nil
}

//
// Create implements TransactionRepository
func (trx *transactionRepository) Create(bill *model.Bill) error {
	result := trx.db.Create(bill).Error
	return result
}

// Delete implements TransactionRepository
func (trx *transactionRepository) Delete(bill *model.Bill) error {
	panic("unimplemented")
}

// ReadAll implements TransactionRepository
func (trx *transactionRepository) ReadAll() ([]model.Bill, error) {
	var trans []model.Bill
	result := trx.db.Find(&trans)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return trans, nil
		} else {
			return trans, err
		}
	}
	return trans, nil
}

// Update implements TransactionRepository
func (trx *transactionRepository) Update(bill *model.Bill, by map[string]interface{}) error {
	panic("unimplemented")
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	repo := new(transactionRepository)
	repo.db = db
	return repo
}
