package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(discount *model.Discount) error
	ReadAll() ([]model.Discount, error)
	// ReadPreloadById(by map[string]interface{}, preload string) (model.TransType, error)
	ReadById(id int) (model.Discount, error)
	// ReadPreloadBy(by map[string]interface{}, preload string) ([]model.TransType, error)
	Update(discount *model.Discount, by map[string]interface{}) error
	// UpdateBy(existingTable *model.TransType) error // untuk yang ber relasi
	Delete(discount *model.Discount) error
}

type discountRepository struct {
	db *gorm.DB
}

// ReadById implements DiscountRepository
func (d *discountRepository) ReadById(id int) (model.Discount, error) {
	var discount model.Discount
	result := d.db.First(&discount, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return discount, nil
		} else {
			return discount, err
		}
	}
	return discount, nil
}

// Create implements DiscountRepository
func (d *discountRepository) Create(discount *model.Discount) error {
	result := d.db.Create(discount).Error
	return result
}

// Delete implements DiscountRepository
func (d *discountRepository) Delete(discount *model.Discount) error {
	result := d.db.Delete(discount).Error
	return result
}

// ReadAll implements DiscountRepository
func (d *discountRepository) ReadAll() ([]model.Discount, error) {
	var discount []model.Discount
	result := d.db.Find(&discount)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return discount, nil
		} else {
			return discount, err
		}
	}
	return discount, nil
}

// Update implements DiscountRepository
func (d *discountRepository) Update(discount *model.Discount, by map[string]interface{}) error {
	result := d.db.Model(discount).Updates(by).Error
	return result
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	repo := new(discountRepository)
	repo.db = db
	return repo
}
