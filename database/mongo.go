package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(" DB Connection Failed:", err)
	}

	Client = client
	log.Println(" Connected to MongoDB")
}

func GetCollection(name string) *mongo.Collection {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")

	return Client.Database(dbName).Collection(name)
}
