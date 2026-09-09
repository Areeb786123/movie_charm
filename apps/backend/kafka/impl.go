package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{
		writer: writer,
	}
}

func (p *KafkaProducer) Publish(
	ctx context.Context,
	topic string,
	key []byte,
	value []byte,
) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	})
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
