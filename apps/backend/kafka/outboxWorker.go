package kafka

import (
	"backend/repository/outbox"
	"context"
	"log"
	"strconv"
)

type OutboxWorker struct {
	outboxRepo outbox.KafkaOutBoxRepo
	producer   Producer
}

func NewOutboxWorker(
	outboxRepo outbox.KafkaOutBoxRepo,
	producer Producer,
) *OutboxWorker {
	return &OutboxWorker{
		outboxRepo: outboxRepo,
		producer:   producer,
	}
}

func (w *OutboxWorker) Process(ctx context.Context) error {

	events, err := w.outboxRepo.GetPendingEvents(ctx, 100)
	if err != nil {
		return err
	}

	for _, event := range events {

		key := []byte(strconv.FormatInt(event.AggregateID, 10))
		err := w.producer.Publish(
			ctx,
			"movie-rated",
			key,
			event.Payload,
		)

		if err != nil {
			log.Printf(
				"failed to publish outbox event %d: %v",
				event.ID,
				err,
			)
			continue
		}

		if err := w.outboxRepo.MarkPublished(ctx, event.ID); err != nil {
			log.Printf(
				"failed to mark outbox event %d as published: %v",
				event.ID,
				err,
			)
		}
	}

	return nil
}
