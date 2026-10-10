package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"delivery-chi-project/cmd/api/handler"
)

func NewRouter(cartHandler *handler.CartHandler, inventoryHandler *handler.InventoryHandler) *chi.Mux {
	r := chi.NewRouter()

	// Enable CORS for Svelte frontend (usually runs on http://localhost:5173)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api", func(r chi.Router) {
		// Cart routes
		r.Route("/cart", func(r chi.Router) {
			r.Post("/", cartHandler.AddToCart)
			r.Get("/", cartHandler.GetCart)
		})

		// Inventory & Product CRUD routes
		r.Route("/inventory", func(r chi.Router) {
			r.Get("/", inventoryHandler.GetInventory)
		})
	})

	return r
}