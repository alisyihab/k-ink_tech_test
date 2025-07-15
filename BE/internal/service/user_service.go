package service

import (
	"context"
	"errors"
	"registration-smart-reward/internal/repo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type userService struct {
	repo       repo.UserRepository
	jwtSecret  []byte
	bcryptCost int
}

func NewUserService(repo repo.UserRepository, jwtSecret string) UserService {
	return &userService{
		repo:       repo,
		jwtSecret:  []byte(jwtSecret),
		bcryptCost: bcrypt.DefaultCost,
	}
}

func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", errors.New("username atau password salah")
	}

	// compare password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("username atau password salah")
	}

	// generate JWT
	claims := jwt.MapClaims{
		"sub": user.ID.Hex(),
		"usr": user.Username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
