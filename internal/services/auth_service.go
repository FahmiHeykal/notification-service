package services

import (
	"errors"
	"notification-service/internal/models"
	"notification-service/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) Register(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}
	return s.userRepo.Create(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.jwtSecret))
}

var ErrInvalidCredentials = errors.New("invalid credentials")
