package rebalancing

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Rebalancer interface {
	MoveIn(ctx context.Context, IDs []string) error
	MoveOut(ctx context.Context, IDs []string) error
}

type serviceCtx struct {
	mongoClientPrimary *mongo.Client
	mongoClientHot     *mongo.Client
}

func New(mongoClientPrimary *mongo.Client, mongoClientHot *mongo.Client) Rebalancer {
	return &serviceCtx{
		mongoClientPrimary: mongoClientPrimary,
		mongoClientHot:     mongoClientHot,
	}
}

func (sc *serviceCtx) MoveIn(ctx context.Context, IDs []string) error {
	return nil
}

func (sc *serviceCtx) MoveOut(ctx context.Context, IDs []string) error {
	return nil
}
