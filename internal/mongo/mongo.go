package mongo

import (
	"context"
	"log"
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
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	client = cl
}
