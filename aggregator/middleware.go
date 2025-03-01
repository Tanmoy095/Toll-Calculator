package main

import (
	"log"
	"time"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	next Aggregator
}

func NewLoggingMiddleware(next Aggregator) *LoggingMiddleware {
	return &LoggingMiddleware{
		next: next,
	}
}

func (l *LoggingMiddleware) AggregateDistance(distance types.Distance) (err error) {
	log.Printf("Processing and inserting distance in the storage: %v", distance)
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{

			"took":     time.Since(start),
			"error":    err,
			"distance": distance,
		}).Info("Processing and inserting distance in the storage:")
	}(time.Now())
	err = l.next.AggregateDistance(distance)
	return
}
