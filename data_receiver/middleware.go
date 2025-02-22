package main

import (
	"time"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Dataproducer
}

func NewLogMiddleware(next Dataproducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Latitude,
			"long":  data.Longitude,
			"took":  time.Since(start),
		}).Info("ProduceData to Kafka server")
	}(time.Now())

	return l.next.ProduceData(data)
}
