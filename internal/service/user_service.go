package service

import (
	"delivery-chi-project/internal/db"
	"delivery-chi-project/internal/models"
)

type UserService struct {
	dbManager *db.DBManager
}

func NewUserService(dbManager *db.DBManager) *UserService {
	return &UserService{dbManager: dbManager}
}

func (s *UserService) GetUserByID(id string) (models.User, bool) {
	data, err := s.dbManager.Read()
	if err != nil {
		return models.User{}, false
	}

	for _, u := range data.Users {
		if u.ID == id {
			return u, true
		}
	}
	return models.User{}, false
}