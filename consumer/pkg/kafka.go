package pkg

import (
	"context"
	"crypto/tls"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"log"
)

type KafkaReader struct {
	kr *kafka.Reader
}

func NewKafkaReader(topic, username, password, url string) *KafkaReader {
	mechanism, err := scram.Mechanism(scram.SHA256, username, password)
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{url},
		Topic:   topic,
		Dialer:  dialer,
	})

	return &KafkaReader{kr: reader}
}

func (kr *KafkaReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return kr.kr.ReadMessage(ctx)
}

func (kr *KafkaReader) Close() error {
	return kr.kr.Close()
}
