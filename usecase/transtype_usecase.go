package usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type TransTypeUseCase interface {
	CreateNewTransType(transType *model.TransType) error
	FindAllTransType() ([]model.TransType, error)
	FindById(id int) (model.TransType, error)
	UpdateTransType(transType *model.TransType, by map[string]interface{}) error
	DeleteTransType(transType *model.TransType) error
}

type transTypeUseCase struct {
	repo repository.TransTypeRepository
}

// FindById implements TransTypeUseCase
func (tr *transTypeUseCase) FindById(id int) (model.TransType, error) {
	return tr.repo.ReadById(id)
}

// CreateNewTransType implements TransTypeUseCase
func (tr *transTypeUseCase) CreateNewTransType(transType *model.TransType) error {
	return tr.repo.Create(transType)
}

// DeleteTransType implements TransTypeUseCase
func (tr *transTypeUseCase) DeleteTransType(transType *model.TransType) error {
	return tr.repo.Delete(transType)
}

// FindAllTransType implements TransTypeUseCase
func (tr *transTypeUseCase) FindAllTransType() ([]model.TransType, error) {
	return tr.repo.ReadAll()
}

// UpdateTransType implements TransTypeUseCase
func (tr *transTypeUseCase) UpdateTransType(transType *model.TransType, by map[string]interface{}) error {
	return tr.repo.Update(transType, by)
}

func NewTransTypeUseCase(repo repository.TransTypeRepository) TransTypeUseCase {
	usc := new(transTypeUseCase)
	usc.repo = repo
	return usc
}
