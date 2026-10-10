package response

import "delivery-chi-project/internal/models"

type CartResponse struct {
	UserID string            `json:"user_id"`
	Items  []models.CartItem `json:"items"`
	Total  float64           `json:"total"`
}