package service

import (
	"delivery-chi-project/internal/db"
	"delivery-chi-project/internal/models"
)

type ProductService struct {
	dbManager *db.DBManager
}

func NewProductService(dbManager *db.DBManager) *ProductService {
	return &ProductService{dbManager: dbManager}
}

func (s *ProductService) GetAllProducts() []models.Product {
	data, err := s.dbManager.Read()
	if err != nil {
		return []models.Product{}
	}
	return data.Products
}

func (s *ProductService) GetProductByID(id string) (models.Product, bool) {
	data, err := s.dbManager.Read()
	if err != nil {
		return models.Product{}, false
	}

	for _, p := range data.Products {
		if p.ID == id {
			return p, true
		}
	}
	return models.Product{}, false
}