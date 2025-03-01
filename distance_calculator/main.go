package main

import (
	"log"

	"github.com/Tanmoy095/Toll-Calculator.git/aggregator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))

	if err != nil {
		log.Fatal(err)

	}
	kafkaConsumer.Start()

}
