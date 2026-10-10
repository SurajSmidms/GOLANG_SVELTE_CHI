package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"delivery-chi-project/cmd/api/handler"
	"delivery-chi-project/internal/db"
	"delivery-chi-project/internal/service"
)

func TestServerRouting(t *testing.T) {
	// Create temporary db file for server tests
	tmpDir := t.TempDir()
	tmpDbPath := filepath.Join(tmpDir, "db.json")
	initialContent := `{
		"products": [
			{"id": "prod-1", "name": "Cheeseburger", "price": 8.99}
		],
		"carts": [],
		"users": []
	}`
	_ = os.WriteFile(tmpDbPath, []byte(initialContent), 0644)

	dbManager := db.NewDBManager(tmpDbPath)
	ps := service.NewProductService(dbManager)
	cs := service.NewCartService(dbManager, ps)

	cartHandler := handler.NewCartHandler(cs)
	inventoryHandler := handler.NewInventoryHandler(ps)

	srv := NewServer("8080", cartHandler, inventoryHandler)

	ts := httptest.NewServer(srv.httpServer.Handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/inventory")
	if err != nil {
		t.Fatalf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}