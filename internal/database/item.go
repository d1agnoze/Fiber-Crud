package database

import (
	"fmt"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name    string
	Price   float64
	InStock bool
}

type CreateOrUpdateItemInput struct {
	Name    string  `validate:"required,min=1,max=100"`
	Price   float64 `validate:"required,min=1"`
	InStock bool    `validate:"required,boolean"`
}

type ItemService interface {
	CreateItem(*CreateOrUpdateItemInput) (uint, error)
	DeleteItem(uint) error
	GetItem(uint) (Item, error)
	GetAllItems() ([]Item, error)
	UpdateItem(uint, *CreateOrUpdateItemInput) error
}

func (s *service) CreateItem(input *CreateOrUpdateItemInput) (uint, error) {
	item := &Item{
		Name:    input.Name,
		Price:   input.Price,
		InStock: input.InStock,
	}

	if err := s.db.Create(&item).Error; err != nil {
		return 0, err
	}

	return item.ID, nil
}

func (s *service) DeleteItem(id uint) error {
	if err := s.db.Delete(&Item{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (s *service) GetItem(id uint) (Item, error) {
	var item Item

	if err := s.db.First(&item, id).Error; err != nil {
		return Item{}, err
	}

	return item, nil
}

func (s *service) GetAllItems() ([]Item, error) {
	var items []Item

	if err := s.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (s *service) UpdateItem(id uint, input *CreateOrUpdateItemInput) error {
	var item Item

	if err := s.db.First(&item, id).Error; err != nil {
		return fmt.Errorf("Item not found, ID: %d", id)
	}

	if err := s.db.Model(&item).Updates(*input).Error; err != nil {
		return fmt.Errorf("Failed to update item, ID: %d", id)
	}

	return nil
}
