package database

import (
	"server-bot/internal/models"
)

func (s *service) GetUser(id uint) (*models.User, error) {
	var user models.User
	if result := s.db.Preload("Addresses").First(&user, id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *service) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if result := s.db.Preload("Addresses").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
