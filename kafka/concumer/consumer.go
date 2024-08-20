package concumer

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type CafkaConcumer interface {
	ConsumerMasagae(handler func(massage []byte)) error
	Close() error
}

type KafkaConsumerImpl struct {
	reader kafka.Reader
	logger *slog.Logger
}

func NewKafkaConsumer(brokers []string, topic string, logger *slog.Logger) CafkaConcumer {
	return &KafkaConsumerImpl{
		reader: *kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		logger: logger,
	}
}

func (c *KafkaConsumerImpl) ConsumerMasagae(handler func(message []byte)) error {
	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			c.logger.Error("Error reading message", "error", err)
			continue
		}
		handler(msg.Value)
	}
}

func (c *KafkaConsumerImpl) Close() error {
	return c.reader.Close()
}
