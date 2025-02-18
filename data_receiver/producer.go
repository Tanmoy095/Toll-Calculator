package main

import (
	"encoding/json"
	"fmt"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Dataproducer interface {
	ProduceData(types.ObuData) error
}
type kafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (*kafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})

	if err != nil {
		return nil, err
	}
	// Start another goroutine to check if we have delivered the data.
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("DeLivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return &kafkaProducer{
		producer: p,
	}, nil
}

//kafka producer is going to implement the interface DataProducer

func (p *kafkaProducer) ProduceData(data types.ObuData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kafkaTopic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)

}
