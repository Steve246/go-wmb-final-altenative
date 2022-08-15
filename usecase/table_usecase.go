package usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
	"log"
)

type TableUseCase interface {
	CreateNewTable(table *model.Table) error
	FindAll() ([]model.Table, error)
	FindById(id int) (model.Table, error)
	CheckIfAvail(id int) (status bool)
	UpdateTable(table *model.Table, by map[string]interface{}) error
	DeleteTable(table *model.Table) error
}

type tableUseCase struct {
	repo repository.TableRepository
}

// FindById implements TableUseCase
func (t *tableUseCase) FindById(id int) (model.Table, error) {
	table, err := t.repo.ReadById(id)
	if err != nil {
		log.Println(err)
		return table, err
	}
	return table, nil
}

// CheckIfAvail implements TableUseCase
func (t *tableUseCase) CheckIfAvail(id int) (status bool) {
	table, err := t.repo.ReadById(id)
	if err != nil {
		log.Println(err)
	}
	return table.IsAvailable
}

// DeleteTable implements TableUseCase
func (t *tableUseCase) DeleteTable(table *model.Table) error {
	return t.repo.Delete(table)
}

// Update implements TableUseCase
func (t *tableUseCase) UpdateTable(table *model.Table, by map[string]interface{}) error {
	return t.repo.Update(table, by)
}

// FindAll implements TableUseCase
func (t *tableUseCase) FindAll() ([]model.Table, error) {
	return t.repo.ReadAll()
}

// Create implements TableUseCase
func (t *tableUseCase) CreateNewTable(table *model.Table) error {
	return t.repo.Create(table)
}

func NewTableUseCase(repo repository.TableRepository) TableUseCase {
	usc := new(tableUseCase)
	usc.repo = repo
	return usc
}
