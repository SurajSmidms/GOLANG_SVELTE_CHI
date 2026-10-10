package models

type GroceryProduct struct {
	Product
	WeightInGrams int `json:"weight_in_grams"`
	IsOrganic     bool `json:"is_organic"`
}