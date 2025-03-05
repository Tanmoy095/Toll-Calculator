// producer.go - Kafka producer logic

package main

import (
	"encoding/json"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// DataProducer interface defines the method to send data to Kafka.
type DataProducer interface {
	ProduceData(types.OBUData) error
}

// KafkaProducer struct holds Kafka producer instance and topic name.
type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

// NewKafkaProducer initializes a Kafka producer for a given topic.
func NewKafkaProducer(topic string) (DataProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}

	// Goroutine to monitor delivery reports
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					// fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					// fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

// ProduceData sends OBU data as a Kafka message.
func (p *KafkaProducer) ProduceData(data types.OBUData) error {
	// Marshal OBU data to JSON
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Send message to Kafka
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)
}
