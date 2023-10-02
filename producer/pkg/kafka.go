package pkg

import (
	"context"
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"log"
)

type KafkaWriter struct {
	kw *kafka.Writer
}

func NewKafkaWriter(topic, username, password, url string) *KafkaWriter {
	mechanism, err := scram.Mechanism(scram.SHA256, username, password)
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{url},
		Topic:   topic,
		Dialer:  dialer,
	})

	return &KafkaWriter{kw: writer}
}

func (kw *KafkaWriter) WriteMessages(ctx context.Context, messages ...kafka.Message) error {
	return kw.kw.WriteMessages(ctx, messages...)
}

func (kw *KafkaWriter) Close() error {
	return kw.kw.Close()
}
