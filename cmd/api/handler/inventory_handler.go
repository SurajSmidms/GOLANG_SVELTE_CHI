package handler

import (
	"net/http"
	"delivery-chi-project/internal/service"
)

type InventoryHandler struct {
	productService *service.ProductService
}

func NewInventoryHandler(ps *service.ProductService) *InventoryHandler {
	return &InventoryHandler{productService: ps}
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	products := h.productService.GetAllProducts()
	RespondJSON(w, http.StatusOK, products)
}