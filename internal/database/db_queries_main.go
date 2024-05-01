package database

import (
	"log"
	"server-bot/internal/models"
)

func (s *service) Health() map[string]string {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	//err := s.db.Ping(ctx, nil)
	//if err != nil {
	//	log.Fatalf(fmt.Sprintf("db down: %v", err))
	//}
	//
	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetAllProducts() ([]*models.Product, error) {
	var prods []*models.Product

	result := s.db.Find(&prods)
	if result.Error != nil {
		log.Printf("Error fetching products: %v\n", result.Error)
		return nil, result.Error
	}

	return prods, nil
}

func (s *service) GetProduct(id uint) (*models.Product, error) {
	var prod models.Product

	result := s.db.First(&prod, id)
	if result.Error != nil {
		log.Printf("Error fetching product with ID %d: %v\n", id, result.Error)
		return nil, result.Error
	}

	return &prod, nil
}

func (s *service) AddProduct(prod *models.Product) error {
	result := s.db.Create(prod)
	if result.Error != nil {
		log.Printf("Error adding product: %v\n", result.Error)
		return result.Error
	}

	return nil
}
