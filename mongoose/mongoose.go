package mongoose

import (
	"context"
	"log"
	"time"

	"example.com/rest-api/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func InitMongoose() {
	MONGODB_URI := env.Get("MONGODB_URI")
	MONGODB_DB_NAME := env.Get("MONGODB_DB_NAME")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(MONGODB_URI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	errPing := client.Database(MONGODB_DB_NAME).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	if errPing != nil {
		log.Fatal(errPing)
	}

	MongoDB = client.Database(MONGODB_DB_NAME)
	log.Print("successfully connected to MongoDB!")

}
