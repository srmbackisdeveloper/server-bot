package models

import (
	"gorm.io/gorm"
	"time"
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

type User struct {
	gorm.Model
	Email            string    `json:"email" gorm:"unique;index"`
	Name             string    `json:"name"`
	Addresses        []Address `json:"addresses" gorm:"foreignKey:UserId"`
	IsActive         bool      `json:"isActive"`
	VerificationCode string    `json:"verificationCode"`
	CodeValidUntil   time.Time `json:"codeValidUntil"`

	//
	Token           string    `json:"token"`
	TokenValidUntil time.Time `json:"tokenValidUntil"`
}

type Address struct {
	gorm.Model
	UserId   uint   `json:"userId" gorm:"index"` // Index for better lookup performance
	Address  string `json:"address"`
	IsActive bool   `json:"isActive"`
}
