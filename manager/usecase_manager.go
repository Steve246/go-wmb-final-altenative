package manager

import "livecode-wmb-2/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
	TableUseCase() usecase.TableUseCase
	TransType() usecase.TransTypeUseCase
	Discount() usecase.DiscountUseCase
	Customer() usecase.CustomerUseCase
	Transaction() usecase.TransactionUseCase
	LopeiChekBalance() usecase.LopeiUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

// CheckBalanceUseCase implements UseCaseManager
func (u *useCaseManager) LopeiChekBalance() usecase.LopeiUseCase {
	return usecase.NewCheckBalanceUseCase(u.repoManager.LopeiRepo())
}

// Transaction implements UseCaseManager
func (u *useCaseManager) Transaction() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repoManager.Transaction())
}

// Customer implements UseCaseManager
func (u *useCaseManager) Customer() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.Customer())
}

// Discount implements UseCaseManager
func (u *useCaseManager) Discount() usecase.DiscountUseCase {
	return usecase.NewDiscountUseCase(u.repoManager.Discount())
}

// TransType implements UseCaseManager
func (u *useCaseManager) TransType() usecase.TransTypeUseCase {
	return usecase.NewTransTypeUseCase(u.repoManager.TransType())
}

// TableUseCase implements UseCaseManager
func (u *useCaseManager) TableUseCase() usecase.TableUseCase {
	return usecase.NewTableUseCase(u.repoManager.TableRepo())
}

// ListProductUseCase implements UseCaseManager
func (u *useCaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(u.repoManager.MenuRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
