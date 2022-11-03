package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	mongoClientPrimary, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoPrimaryShardConnectionString))
	if err != nil {
		panic(err)
	}

	mongoClientHot, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDBHotShardConnectionString))
	if err != nil {
		panic(err)
	}
}
