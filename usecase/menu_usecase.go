package usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type MenuUseCase interface {
	InsertMenu(menu *model.Menu) error
	ShowAllMenu(preload string) ([]model.Menu, error)
	FindMenuById(by map[string]interface{}, preload string) (model.Menu, error)
	FindById(id string) (model.Menu, error)
	UpdateMenu(menu *model.Menu, by map[string]interface{}) error
	UpdatePrice(existingMenu *model.Menu) error // untuk yang ber relasi
	DeleteMenu(menu *model.Menu) error
}

type menuUseCase struct {
	repo repository.MenuRepository
}

// FindById implements MenuUseCase
func (m *menuUseCase) FindById(id string) (model.Menu, error) {
	return m.repo.ReadById(id)
}

// DeleteMenu implements MenuUseCase
func (m *menuUseCase) DeleteMenu(menu *model.Menu) error {
	return m.repo.Delete(menu)
}

// FindMenuById implements MenuUseCase
func (m *menuUseCase) FindMenuById(by map[string]interface{}, preload string) (model.Menu, error) {
	return m.repo.ReadPreloadById(by, preload)

}

// InsertMenu implements MenuUseCase
func (m *menuUseCase) InsertMenu(menu *model.Menu) error {
	return m.repo.Create(menu)

}

// ShowAllMenu implements MenuUseCase
func (m *menuUseCase) ShowAllMenu(preload string) ([]model.Menu, error) {
	return m.repo.ReadPreloadAll(preload)
}

// UpdateMenu implements MenuUseCase
func (m *menuUseCase) UpdateMenu(menu *model.Menu, by map[string]interface{}) error {
	result := m.repo.Update(menu, by)
	return result
}

// UpdatePrice implements MenuUseCase
func (m *menuUseCase) UpdatePrice(existingMenu *model.Menu) error {
	return m.repo.UpdateBy(existingMenu)
}

func NewMenuUseCase(repo repository.MenuRepository) MenuUseCase {
	usc := new(menuUseCase)
	usc.repo = repo
	return usc
}
