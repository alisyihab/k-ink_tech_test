package main

import (
	"context"
	"log"
	"os"
	"time"

	"registration-smart-reward/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Gagal konek MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("registration_smart_reward")
	col := db.Collection("paket")

	pakets := []interface{}{
		model.Paket{
			Wilayah:   "WIL.A",
			Tipe:      "Fast Track",
			Level:     "Silver",
			Nama:      "Paket Silver A",
			Harga:     100000,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		model.Paket{
			Wilayah:   "WIL.A",
			Tipe:      "Smart Reward",
			Level:     "Gold",
			Nama:      "Paket Gold A",
			Harga:     200000,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		model.Paket{
			Wilayah:   "WIL.B",
			Tipe:      "Fast Track",
			Level:     "Platinum",
			Nama:      "Paket Platinum B",
			Harga:     500000,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if _, err := col.InsertMany(context.Background(), pakets); err != nil {
		log.Fatal("Gagal seeding paket:", err)
	}

	log.Println("Seeder paket berhasil dijalankan")
}
