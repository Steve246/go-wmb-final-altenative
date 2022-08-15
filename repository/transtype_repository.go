package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type TransTypeRepository interface {
	Create(transType *model.TransType) error
	ReadAll() ([]model.TransType, error)
	ReadById(id int) (model.TransType, error)
	Update(transType *model.TransType, by map[string]interface{}) error
	Delete(transType *model.TransType) error
}

type transTypeRepository struct {
	db *gorm.DB
}

// ReadById implements TransTypeRepository
func (t *transTypeRepository) ReadById(id int) (model.TransType, error) {
	var transType model.TransType
	result := t.db.First(&transType, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transType, nil
		} else {
			return transType, err
		}
	}
	return transType, nil
}

// Create implements TransTypeRepository
func (tr *transTypeRepository) Create(transType *model.TransType) error {
	result := tr.db.Create(transType).Error
	return result
}

// Delete implements TransTypeRepository
func (tr *transTypeRepository) Delete(transType *model.TransType) error {
	result := tr.db.Delete(transType).Error
	return result
}

// ReadAll implements TransTypeRepository
func (tr *transTypeRepository) ReadAll() ([]model.TransType, error) {
	var transType []model.TransType
	result := tr.db.Find(&transType)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transType, nil
		} else {
			return transType, err
		}
	}
	return transType, nil
}

// Update implements TransTypeRepository
func (tr *transTypeRepository) Update(transType *model.TransType, by map[string]interface{}) error {
	result := tr.db.Model(transType).Updates(by).Error
	return result
}

func NewTransTypeRepository(db *gorm.DB) TransTypeRepository {
	repo := new(transTypeRepository)
	repo.db = db
	return repo
}
