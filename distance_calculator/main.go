package main

import "log"

const kafkaTopic = "obudata"

func main() {
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic)
	if err != nil {
		log.Fatal(err)

	}
	kafkaConsumer.Start()

}
