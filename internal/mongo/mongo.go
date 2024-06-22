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
	client *mongo.Client
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
}
