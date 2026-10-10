package service

import "delivery-chi-project/internal/models"

var MockProducts = []models.Product{
	{ID: "prod-1", Name: "Cheeseburger", Price: 8.99},
	{ID: "prod-2", Name: "Margherita Pizza", Price: 12.49},
	{ID: "prod-3", Name: "Organic Apples (1kg)", Price: 4.50},
}