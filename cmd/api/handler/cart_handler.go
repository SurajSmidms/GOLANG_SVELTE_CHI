package handler

import (
	"encoding/json"
	"net/http"
	"delivery-chi-project/internal/dto/request"
	"delivery-chi-project/internal/service"
)

type CartHandler struct {
	cartService *service.CartService
}

func NewCartHandler(cs *service.CartService) *CartHandler {
	return &CartHandler{cartService: cs}
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req request.AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	cart, err := h.cartService.AddToCart(req.UserID, req.ProductID, req.Quantity)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cart)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		RespondError(w, http.StatusBadRequest, "Missing user_id query parameter")
		return
	}

	cart, err := h.cartService.GetCart(userID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, cart)
}