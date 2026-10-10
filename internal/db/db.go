package db

import (
	"encoding/json"
	"os"
	"sync"
	"delivery-chi-project/internal/models"
)

type Database struct {
	Products []models.Product `json:"products"`
	Carts    []models.Cart    `json:"carts"`
	Users    []models.User    `json:"users"`
}

type DBManager struct {
	mu       sync.Mutex
	filePath string
}

func NewDBManager(filePath string) *DBManager {
	return &DBManager{filePath: filePath}
}

func (db *DBManager) Read() (*Database, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.Open(db.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data Database
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (db *DBManager) Write(data *Database) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.Create(db.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}