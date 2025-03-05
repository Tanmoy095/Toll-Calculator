// middleware.go - Logging middleware for Kafka producer

package main

import (
	"time"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/sirupsen/logrus"
)

// LogMiddleware wraps a producer with logging functionality.
type LogMiddleware struct {
	next DataProducer
}

// NewLogMiddleware creates a new logging middleware.
func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

// ProduceData logs metadata before sending data to Kafka.
func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Latitude,
			"long":  data.Longitude,
			"took":  time.Since(start), // Time taken to process message
		}).Info("Producing to Kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
