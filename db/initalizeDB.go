package db

import (
    "context"
    "fmt"
    "log"
	"time"

    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func DbInitalize()*mongo.Client{
	mongoconn1 := "mongodb+srv://eyobmamo25:Eyob%401993@cluster0.ljc5e.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
    clientOptions := options.Client().ApplyURI(mongoconn1)


	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client



}

