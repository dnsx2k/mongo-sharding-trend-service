package main

import (
	"context"

	"github.com/dnsx2k/mongo-sharding-trend-service/cmd/worker/trendmsghandler"
	"github.com/dnsx2k/mongo-sharding-trend-service/pkg/lookupclient"
	"github.com/dnsx2k/mongo-sharding-trend-service/pkg/rebalancing"
	amqp "github.com/rabbitmq/amqp091-go"
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

	//"amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(cfg.AMQPConnectionString)
	if err != nil {
		panic(err)
	}

	lookupClient := lookupclient.New(cfg.LookupServiceBaseURL)
	rebalancer := rebalancing.New(mongoClientPrimary, mongoClientHot, lookupClient)
	msgHandler := trendmsghandler.New(conn, rebalancer, "trendapp.q.trendUpdate")
	go func() {
		_ = msgHandler.Process(context.Background())
	}()
}
