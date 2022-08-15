package usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
	"log"
)

type CustomerUseCase interface {
	Registration(customer *model.Customer) error
	GetAllCustomer(preload string) ([]model.Customer, error)
	ActivationMemberForExistingCustomerAndAddDiscount(id int, description string, pct int) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// GetAllCustomer implements CustomerUseCase
func (c *customerUseCase) GetAllCustomer(prelaod string) ([]model.Customer, error) {
	return c.repo.ReadAllPreload(prelaod)
}

// ActivationMemberForExistingCustomerAndAddDiscount implements CustomerUseCase
func (c *customerUseCase) ActivationMemberForExistingCustomerAndAddDiscount(id int, description string, pct int) error {
	cust, err := c.repo.ReadById(id)
	if err != nil {
		return err
	}

	err = c.repo.Update(&cust, map[string]interface{}{
		"is_member": "true"})

	if err != nil {
		log.Println(err)
		return err
	}
	cust.Discount = []*model.Discount{{
		Description: description,
		Pct:         pct,
	},
	}
	err = c.repo.UpdateBy(&cust)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Registration implements CustomerUseCase
func (c *customerUseCase) Registration(customer *model.Customer) error {
	return c.repo.Create(customer)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	usc := new(customerUseCase)
	usc.repo = repo
	return usc
}
