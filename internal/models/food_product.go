package models

type FoodProduct struct {
	Product
	Calories int `json:"calories"`
	IsVeg    bool `json:"is_veg"`
}