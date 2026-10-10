package main

import (
	"log"

	"delivery-chi-project/cmd/api/handler"
	"delivery-chi-project/cmd/server"
	"delivery-chi-project/internal/db"
	"delivery-chi-project/internal/service"
)

func main() {
	// Initialize JSON File Database Manager
	dbManager := db.NewDBManager("db.json")

	// Initialize Services with DB manager
	productService := service.NewProductService(dbManager)
	cartService := service.NewCartService(dbManager, productService)

	// Initialize Handlers
	cartHandler := handler.NewCartHandler(cartService)
	inventoryHandler := handler.NewInventoryHandler(productService)

	// Initialize Server with Chi router
	srv := server.NewServer("8080", cartHandler, inventoryHandler)

	log.Println("Starting server with Chi router and JSON file DB on port 8080...")
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}