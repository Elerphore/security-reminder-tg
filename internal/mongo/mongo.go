package mongo

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client              *mongo.Client
	user_collection     *mongo.Collection
	days_collection     *mongo.Collection
	messages_collection *mongo.Collection
	ctx                 = context.Background()
)

var defaultLog = log.Default()

func MongoDBConncetion(wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.WithoutCancel(context.Background())
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URI")))

	if err != nil {
		log.Default().Fatalln(err)
	}

	client = cl

	user_collection = client.Database("main").Collection("user")
	days_collection = client.Database("main").Collection("days")
	messages_collection = client.Database("main").Collection("messages")
}
