package producer

import (
	"encoding/json"
	"time"
	"github.com/Shopify/sarama"
	)

type Producer interface {
	Send(message Message) error
}

type producer struct {
	prod sarama.SyncProducer
	topic string
}

func NewProducer(addr []string, topic string) (*producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		prod: syncProducer,
		topic: topic,
	}, nil
}

func (p *producer) Send(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Value:     sarama.StringEncoder(bytes),
		Timestamp: time.Time{},
	}
	_, _, err = p.prod.SendMessage(msg)
	return err
}
