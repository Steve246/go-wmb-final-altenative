package manager

import (
	"livecode-wmb-2/repository"
)

type RepositoryManager interface {
	MenuRepo() repository.MenuRepository
	TableRepo() repository.TableRepository
	TransType() repository.TransTypeRepository
	Discount() repository.DiscountRepository
	Customer() repository.CustomerRepository
	Transaction() repository.TransactionRepository
	LopeiRepo() repository.LopeiRepository
}

type repositoryManager struct {
	infra Infra
}

// CustomerRepo implements RepositoryManager
func (r *repositoryManager) LopeiRepo() repository.LopeiRepository {
	return repository.NewLopeiRepository(r.infra.lopeiClientConn())
}

// Transaction implements RepositoryManager
func (r *repositoryManager) Transaction() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.SqlDb())
}

// Customer implements RepositoryManager
func (r *repositoryManager) Customer() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

// Discount implements RepositoryManager
func (r *repositoryManager) Discount() repository.DiscountRepository {
	return repository.NewDiscountRepository(r.infra.SqlDb())
}

// TransType implements RepositoryManager
func (r *repositoryManager) TransType() repository.TransTypeRepository {
	return repository.NewTransTypeRepository(r.infra.SqlDb())
}

// TableRepo implements RepositoryManager
func (r *repositoryManager) TableRepo() repository.TableRepository {
	return repository.NewTableRepository(r.infra.SqlDb())
}

// ProductRepo implements RepositoryManager
func (r *repositoryManager) MenuRepo() repository.MenuRepository {
	return repository.NewMenuRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
