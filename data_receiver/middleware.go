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
func (l *LogMiddleware) ProduceData(data types.ObuData) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("ProduceData to kafka server")

	}(time.Now())

	return nil
}
