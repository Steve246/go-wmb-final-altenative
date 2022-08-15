package usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type DiscountUseCase interface {
	CreateNewDiscount(discount *model.Discount) error
	FindAllDiscount() ([]model.Discount, error)
	FindById(id int) (model.Discount, error)
	UpdateDiscount(discount *model.Discount, by map[string]interface{}) error
	DeleteDiscount(discount *model.Discount) error
}

type discountUseCase struct {
	repo repository.DiscountRepository
}

// FindById implements DiscountUseCase
func (d *discountUseCase) FindById(id int) (model.Discount, error) {
	return d.repo.ReadById(id)
}

// CreateNewDiscount implements DiscountUseCase
func (d *discountUseCase) CreateNewDiscount(discount *model.Discount) error {
	return d.repo.Create(discount)
}

// DeleteTable implements DiscountUseCase
func (d *discountUseCase) DeleteDiscount(discount *model.Discount) error {
	return d.repo.Delete(discount)
}

// FindAllDiscount implements DiscountUseCase
func (d *discountUseCase) FindAllDiscount() ([]model.Discount, error) {
	return d.repo.ReadAll()
}

// UpdateTable implements DiscountUseCase
func (d *discountUseCase) UpdateDiscount(discount *model.Discount, by map[string]interface{}) error {
	return d.repo.Update(discount, by)
}

func NewDiscountUseCase(repo repository.DiscountRepository) DiscountUseCase {
	usc := new(discountUseCase)
	usc.repo = repo
	return usc
}
