package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *model.Menu) error
	ReadPreloadAll(preload string) ([]model.Menu, error)
	ReadPreloadById(by map[string]interface{}, preload string) (model.Menu, error)
	ReadById(id string) (model.Menu, error)
	ReadPreloadBy(by map[string]interface{}, preload string) ([]model.Menu, error)
	Update(menu *model.Menu, by map[string]interface{}) error
	UpdateBy(existingMenu *model.Menu) error // untuk yang ber relasi
	Delete(menu *model.Menu) error
}

type menuRepository struct {
	db *gorm.DB
}

// ReadById implements MenuRepository
func (m *menuRepository) ReadById(id string) (model.Menu, error) {
	var customer model.Menu
	result := m.db.First(&customer, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

// Create implements MenuRepository
func (m *menuRepository) Create(menu *model.Menu) error {
	result := m.db.Create(menu).Error
	return result
}

// Delete implements MenuRepository
func (m *menuRepository) Delete(menu *model.Menu) error {
	result := m.db.Delete(menu).Error
	return result
}

// ReadPreloadAll implements MenuRepository
func (m *menuRepository) ReadPreloadAll(preload string) ([]model.Menu, error) {
	var menu []model.Menu
	result := m.db.Preload(preload).Find(&menu)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, nil
		} else {
			return menu, err
		}
	}
	return menu, nil
}

// ReadPreloadBy implements MenuRepository
func (m *menuRepository) ReadPreloadBy(by map[string]interface{}, preload string) ([]model.Menu, error) {
	var menu []model.Menu
	result := m.db.Preload(preload).Where(by).Find(&menu)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, nil
		} else {
			return menu, err
		}
	}
	return menu, nil
}

// ReadPreloadById implements MenuRepository
func (m *menuRepository) ReadPreloadById(by map[string]interface{}, preload string) (model.Menu, error) {
	var menu model.Menu
	result := m.db.Preload(preload).Where(by).First(&menu)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, nil
		} else {
			return menu, err
		}
	}
	return menu, nil
}

// Update implements MenuRepository
func (m *menuRepository) Update(menu *model.Menu, by map[string]interface{}) error {
	result := m.db.Model(menu).Updates(by).Error
	return result
}

// UpdateBy implements MenuRepository
func (m *menuRepository) UpdateBy(existingMenu *model.Menu) error {
	result := m.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&existingMenu).Error
	return result
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo
}
