package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type TableRepository interface {
	Create(table *model.Table) error
	ReadAll() ([]model.Table, error)
	ReadPreloadById(by map[string]interface{}, preload string) (model.Table, error)
	ReadById(id int) (model.Table, error)
	ReadPreloadBy(by map[string]interface{}, preload string) ([]model.Table, error)
	Update(table *model.Table, by map[string]interface{}) error
	UpdateBy(existingTable *model.Table) error // untuk yang ber relasi
	Delete(table *model.Table) error
}

type tableRepository struct {
	db *gorm.DB
}

// Create implements TableRepository
func (t *tableRepository) Create(table *model.Table) error {
	result := t.db.Create(table).Error
	return result
}

// Delete implements TableRepository
func (t *tableRepository) Delete(table *model.Table) error {
	result := t.db.Delete(table).Error
	return result
}

// ReadById implements TableRepository
func (t *tableRepository) ReadById(id int) (model.Table, error) {
	var table model.Table
	result := t.db.First(&table, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return table, nil
		} else {
			return table, err
		}
	}
	return table, nil
}

// ReadPreloadAll implements TableRepository
func (t *tableRepository) ReadAll() ([]model.Table, error) {
	var table []model.Table
	result := t.db.Find(&table)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return table, nil
		} else {
			return table, err
		}
	}
	return table, nil
}

// ReadPreloadBy implements TableRepository
func (*tableRepository) ReadPreloadBy(by map[string]interface{}, preload string) ([]model.Table, error) {
	panic("unimplemented")
}

// ReadPreloadById implements TableRepository
func (*tableRepository) ReadPreloadById(by map[string]interface{}, preload string) (model.Table, error) {
	panic("unimplemented")
}

// Update implements TableRepository
func (t *tableRepository) Update(table *model.Table, by map[string]interface{}) error {
	result := t.db.Model(table).Updates(by).Error
	return result
}

// UpdateBy implements TableRepository
func (*tableRepository) UpdateBy(existingTable *model.Table) error {
	panic("unimplemented")
}

func NewTableRepository(db *gorm.DB) TableRepository {
	repo := new(tableRepository)
	repo.db = db
	return repo
}
