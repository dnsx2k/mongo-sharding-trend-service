package rebalancing

import (
	"context"

	"github.com/dnsx2k/mongo-sharding-trend-service/pkg/lookupclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Rebalancer interface {
	MoveIn(ctx context.Context, IDs []string) error
	MoveOut(ctx context.Context, IDs []string) error
}

type serviceCtx struct {
	mongoClientPrimary *mongo.Client
	mongoClientHot     *mongo.Client
	lookupClient       *lookupclient.ClientCtx
}

func New(mongoClientPrimary *mongo.Client, mongoClientHot *mongo.Client, lookupClient *lookupclient.ClientCtx) Rebalancer {
	return &serviceCtx{
		mongoClientPrimary: mongoClientPrimary,
		mongoClientHot:     mongoClientHot,
		lookupClient:       lookupClient,
	}
}

func (sc *serviceCtx) MoveIn(ctx context.Context, IDs []string) error {
	collection := sc.mongoClientPrimary.Database("customSharding").Collection("products")
	cursor, err := collection.Find(ctx, bson.M{"id": bson.M{"$in": IDs}})
	if err != nil {
		return err
	}

	var products []interface{}
	if err := cursor.All(ctx, &products); err != nil {
		return err
	}

	collectionHot := sc.mongoClientHot.Database("customSharding").Collection("products")
	_, err = collectionHot.InsertMany(ctx, products)
	if err != nil {
		return err
	}

	if err := sc.lookupClient.SendLookupEntries("hot", IDs); err != nil {
		return err
	}

	_, err = collection.DeleteMany(ctx, bson.M{"id": bson.M{"$in": IDs}})
	if err != nil {
		return err
	}

	return nil
}

func (sc *serviceCtx) MoveOut(ctx context.Context, IDs []string) error {
	collection := sc.mongoClientHot.Database("customSharding").Collection("products")
	cursor, err := collection.Find(ctx, bson.M{"id": bson.M{"$in": IDs}})
	if err != nil {
		return err
	}

	var products []interface{}
	if err := cursor.All(ctx, &products); err != nil {
		return err
	}

	collectionHot := sc.mongoClientPrimary.Database("customSharding").Collection("products")
	_, err = collectionHot.InsertMany(ctx, products)
	if err != nil {
		return err
	}

	if err := sc.lookupClient.DeleteLookupEntries("hot", IDs); err != nil {
		return err
	}

	_, err = collectionHot.DeleteMany(ctx, bson.M{"id": bson.M{"$in": IDs}})
	if err != nil {
		return err
	}

	return nil
}
