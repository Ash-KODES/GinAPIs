package utils

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
    fmt.Println("Initializing database connection...")
    err := connectDB()
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    fmt.Println("Database connection established successfully")
}

func connectDB() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return fmt.Errorf("Error connecting to MongoDB: %w", err)
    }

    // Ping-pong database
    err = client.Ping(ctx, nil)
    if err != nil {
        return fmt.Errorf("Error pinging MongoDB: %w", err)
    }

    Client = client
    return nil
}

func GetCollection(collectionName string) *mongo.Collection {
    if Client == nil {
        log.Fatal("MongoDB client is not initialized")
    }
    return Client.Database("socialapp").Collection(collectionName)
}