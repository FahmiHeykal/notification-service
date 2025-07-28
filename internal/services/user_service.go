package services

import (
	"notification-service/internal/models"
	"notification-service/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(username, password string) error {
	return s.userRepo.Create(&models.User{
		Username: username,
		Password: password,
	})
}
