package usecase

import (
	"livecode-wmb-2/repository"
)

type LopeiUseCase interface {
	GetBalance(lopeId int32) (float32, error)
}

type chackBalanceUseCase struct {
	repo repository.LopeiRepository
}

// GetBalance implements ChackBalanceUseCase
func (c *chackBalanceUseCase) GetBalance(lopeId int32) (float32, error) {
	return c.repo.CheckBalance(lopeId)
}

func NewCheckBalanceUseCase(repo repository.LopeiRepository) LopeiUseCase {
	return &chackBalanceUseCase{
		repo: repo,
	}
}
