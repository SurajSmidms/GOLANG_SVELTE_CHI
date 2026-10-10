package service

import (
	"os"
	"path/filepath"
	"testing"
	"delivery-chi-project/internal/db"
)

func TestCartService_AddToCart(t *testing.T) {
	// Create an isolated temp directory for testing DB operations
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
	ps := NewProductService(dbManager)
	cs := NewCartService(dbManager, ps)

	cart, err := cs.AddToCart("user-1", "prod-1", 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(cart.Items) != 1 {
		t.Fatalf("expected 1 item in cart, got %d", len(cart.Items))
	}

	if cart.Items[0].Quantity != 2 {
		t.Fatalf("expected quantity 2, got %d", cart.Items[0].Quantity)
	}

	expectedTotal := 8.99 * 2
	if cart.Total != expectedTotal {
		t.Fatalf("expected total %f, got %f", expectedTotal, cart.Total)
	}
}