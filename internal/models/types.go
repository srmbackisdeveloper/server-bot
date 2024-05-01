package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	ImageURL    string `json:"imageURL"`

	Weight         float64 `json:"weight"`
	Price          float64 `json:"price"`
	InventoryCount int     `json:"inventoryCount"`

	OnStock  bool `json:"onStock"`
	IsActive bool `json:"isActive"`
}
