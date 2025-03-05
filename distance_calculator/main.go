package main

import (
	"log"

	"github.com/Tanmoy095/Toll-Calculator.git/distance_calculator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	// Creating a client to send data to Aggregator
	aggregatorClient := client.NewClient(aggregatorEndpoint)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, aggregatorClient)

	if err != nil {
		log.Fatal(err)

	}
	kafkaConsumer.Start()

}
