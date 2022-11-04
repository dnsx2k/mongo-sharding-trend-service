package trendmsghandler

import (
	"context"
	"encoding/json"

	"github.com/dnsx2k/mongo-sharding-trend-service/pkg/dto"
	"github.com/dnsx2k/mongo-sharding-trend-service/pkg/rebalancing"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPMessageHandler struct {
	connection *amqp.Connection
	rebalancer rebalancing.Rebalancer
	queue      string
}

func New(connection *amqp.Connection, rebalancer rebalancing.Rebalancer, queue string) *AMQPMessageHandler {
	return &AMQPMessageHandler{connection: connection, rebalancer: rebalancer, queue: queue}
}

func (mHandler *AMQPMessageHandler) Process(ctx context.Context) error {
	ch, err := mHandler.connection.Channel()
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(mHandler.queue, "", false, false, false, false, nil)
	for m := range msgs {
		var op dto.TrendOperation
		if err := json.Unmarshal(m.Body, &op); err != nil {
			//TODO: Handle error
		}

		switch op.Direction {
		case "in":
			if err := mHandler.rebalancer.MoveIn(ctx, op.IDs); err != nil {
				//TODO: Handle error
			}
		case "out":
			if err := mHandler.rebalancer.MoveOut(ctx, op.IDs); err != nil {
				//TODO: Handle error
			}
		}
	}

	return nil
}
