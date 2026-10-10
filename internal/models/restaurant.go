package models

type Restaurant struct {
	Outlet
	CuisineType string `json:"cuisine_type"`
}