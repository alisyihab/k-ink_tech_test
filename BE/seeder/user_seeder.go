package main

import (
	"context"
	"fmt"
	"log"
	"registration-smart-reward/internal/config"
	"registration-smart-reward/internal/model"
	"registration-smart-reward/internal/repo"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load ENV
	if err := godotenv.Load(".env"); err != nil {
		log.Println("ENV not loaded, relying on OS env")
	}

	// Init Mongo
	config.InitMongo()
	ctx := context.Background()

	userRepo := repo.NewUserRepository()

	username := "admin"
	password := "admin123"
	name := "Super Admin"

	// Cek apakah user sudah ada
	existing, _ := userRepo.FindByUsername(ctx, username)
	if existing != nil {
		log.Fatalf("User dengan username '%s' sudah ada!", username)
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Gagal hash password: %v", err)
	}

	user := &model.User{
		Name:      name,
		Username:  username,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := userRepo.Insert(ctx, user); err != nil {
		log.Fatalf("Gagal insert user: %v", err)
	}

	fmt.Println("User admin berhasil dibuat:")
	fmt.Println("   Username:", username)
	fmt.Println("   Password:", password)
}
