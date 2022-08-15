package usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/model/dto"
	"livecode-wmb-2/repository"
)

type TransactionUseCase interface {
	FindAll() ([]model.Bill, error)
	FindById(id int) (model.Bill, error)
	CreateNewTransaction(trx *model.Bill) error
	PrintBillbyId(by map[string]interface{}) (dto.CetakBill, error)
}
type transactionUseCase struct {
	repo repository.TransactionRepository
}

// FindById implements TransactionUseCase
func (trx *transactionUseCase) FindById(id int) (model.Bill, error) {
	return trx.repo.ReadById(id)
}

// FindAll implements TransactionUseCase
func (trx *transactionUseCase) FindAll() ([]model.Bill, error) {
	trans, err := trx.repo.ReadAll()
	return trans, err
}

// PrintBillby implements TransactionUseCase
func (trx *transactionUseCase) PrintBillbyId(by map[string]interface{}) (dto.CetakBill, error) {
	return trx.repo.FindByIDPreload(by)
}

// CreateNewTransaction implements TransactionUseCase
func (trx *transactionUseCase) CreateNewTransaction(transaction *model.Bill) error {
	return trx.repo.Create(transaction)
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	usc := new(transactionUseCase)
	usc.repo = repo
	return usc
}
