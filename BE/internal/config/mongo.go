package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongo() {
	// Load .env file from the correct relative path
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not loaded, relying on environment variables")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	if uri == "" || dbName == "" {
		log.Fatal("Missing MONGO_URI or MONGO_DB in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ MongoDB connection error: %v", err)
	}

	// Test connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	log.Println("Connected to MongoDB:", uri)

	MongoClient = client
	MongoDB = client.Database(dbName)
}

func GetMongoDB() *mongo.Database {
	return MongoDB
}
