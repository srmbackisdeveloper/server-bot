package database

import (
	"gorm.io/gorm"
	"log"
	"server-bot/internal/functionalities"
	"server-bot/internal/models"
	"time"
)

func (s *service) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Printf("Error fetching user: %v\n", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (s *service) RegisterUser(email string) (*models.User, error) {
	user := models.User{
		Email:            email,
		IsActive:         false,
		VerificationCode: functionalities.GenerateCode(),
		CodeValidUntil:   time.Now().Add(15 * time.Minute),
	}
	result := s.db.Create(&user)
	if result.Error != nil {
		log.Printf("Error registering user: %v\n", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (s *service) UpdateUserVerificationCode(user *models.User) error {
	user.VerificationCode = functionalities.GenerateCode()
	user.CodeValidUntil = time.Now().Add(15 * time.Minute)
	result := s.db.Save(user)
	if result.Error != nil {
		log.Printf("Error updating user verification code: %v\n", result.Error)
		return result.Error
	}
	return nil
}

func (s *service) ActivateUser(user *models.User) error {
	result := s.db.Save(user)
	if result.Error != nil {
		log.Printf("Error activating user: %v\n", result.Error)
		return result.Error
	}
	return nil
}
