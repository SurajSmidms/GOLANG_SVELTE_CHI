package models

type GroceryStore struct {
	Outlet
	DeliveryRadiusKm float64 `json:"delivery_radius_km"`
}