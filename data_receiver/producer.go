package main

import (
	"encoding/json"
	"log"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Dataproducer interface {
	ProduceData(types.OBUData) error
}

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (Dataproducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Println("Kafka producer initialization error:", err)
		return nil, err
	}
	return &kafkaProducer{producer: p}, nil
}

func (p *kafkaProducer) ProduceData(data types.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event, 1)
	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kafkaTopic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, deliveryChan)

	if err != nil {
		log.Println("Kafka produce error:", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Println("Kafka delivery failed:", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	log.Println("Kafka message delivered:", m.TopicPartition)
	return nil
}
