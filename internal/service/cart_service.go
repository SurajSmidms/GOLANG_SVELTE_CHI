package service

import (
	"errors"
	"delivery-chi-project/internal/db"
	"delivery-chi-project/internal/models"
)

type CartService struct {
	dbManager      *db.DBManager
	productService *ProductService
}

func NewCartService(dbManager *db.DBManager, ps *ProductService) *CartService {
	return &CartService{
		dbManager:      dbManager,
		productService: ps,
	}
}

func (s *CartService) AddToCart(userID, productID string, quantity int) (*models.Cart, error) {
	product, exists := s.productService.GetProductByID(productID)
	if !exists {
		return nil, errors.New("product not found")
	}

	database, err := s.dbManager.Read()
	if err != nil {
		return nil, err
	}

	var targetCart *models.Cart
	foundCartIndex := -1

	for i, c := range database.Carts {
		if c.UserID == userID {
			targetCart = &database.Carts[i]
			foundCartIndex = i
			break
		}
	}

	if targetCart == nil {
		newCart := models.Cart{UserID: userID, Items: []models.CartItem{}}
		database.Carts = append(database.Carts, newCart)
		targetCart = &database.Carts[len(database.Carts)-1]
	}

	foundItem := false
	for i, item := range targetCart.Items {
		if item.Product.ID == productID {
			targetCart.Items[i].Quantity += quantity
			foundItem = true
			break
		}
	}

	if !foundItem {
		targetCart.Items = append(targetCart.Items, models.CartItem{Product: product, Quantity: quantity})
	}

	s.calculateTotal(targetCart)

	if foundCartIndex != -1 {
		database.Carts[foundCartIndex] = *targetCart
	} else {
		database.Carts[len(database.Carts)-1] = *targetCart
	}

	if err := s.dbManager.Write(database); err != nil {
		return nil, err
	}

	return targetCart, nil
}

func (s *CartService) GetCart(userID string) (*models.Cart, error) {
	database, err := s.dbManager.Read()
	if err != nil {
		return nil, err
	}

	for _, c := range database.Carts {
		if c.UserID == userID {
			return &c, nil
		}
	}

	return &models.Cart{UserID: userID, Items: []models.CartItem{}, Total: 0}, nil
}

func (s *CartService) calculateTotal(cart *models.Cart) {
	var total float64
	for _, item := range cart.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	cart.Total = total
}