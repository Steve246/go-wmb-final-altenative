package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	ReadAll() ([]model.Customer, error)
	ReadAllPreload(preload string) ([]model.Customer, error)
	ReadById(id int) (model.Customer, error)
	ReadBy(by map[string]interface{}) ([]model.Customer, error)
	Update(customer *model.Customer, by map[string]interface{}) error
	UpdateBy(existingCustomer *model.Customer) error // untuk yang ber relasi
	Delete(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

// ReadAllPreload implements CustomerRepository
func (c *customerRepository) ReadAllPreload(preload string) ([]model.Customer, error) {
	var customer []model.Customer
	err := c.db.Model(&customer).Preload(preload).Find(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// UpdateBy implements CustomerRepository
func (c *customerRepository) UpdateBy(existingCustomer *model.Customer) error {
	result := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&existingCustomer).Error
	return result
}

// Create implements CustomerRepository
func (c *customerRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer).Error
	return result
}

// Delete implements CustomerRepository
func (c *customerRepository) Delete(customer *model.Customer) error {
	result := c.db.Delete(customer).Error
	return result
}

// ReadAll implements CustomerRepository
func (c *customerRepository) ReadAll() ([]model.Customer, error) {
	var customer []model.Customer
	result := c.db.Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

// ReadBy implements CustomerRepository
func (c *customerRepository) ReadBy(by map[string]interface{}) ([]model.Customer, error) {
	var customer []model.Customer
	result := c.db.Unscoped().Where(by).Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

// ReadById implements CustomerRepository
func (c *customerRepository) ReadById(id int) (model.Customer, error) {
	var customer model.Customer
	result := c.db.First(&customer, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

// Update implements CustomerRepository
func (c *customerRepository) Update(customer *model.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by).Error
	return result
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
