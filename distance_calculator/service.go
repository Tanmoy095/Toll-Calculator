package main

import (
	"math"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

type CalculatorServicer interface {
	// Add any necessary dependencies here
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	// Add any necessary dependencies here
	prevPoint []float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0

	if len(s.prevPoint) > 0 {
		distance = calculate_Distance(s.prevPoint[0], s.prevPoint[1], data.Latitude, data.Longitude)

	}
	s.prevPoint = []float64{data.Latitude, data.Longitude}
	return distance, nil
}

func calculate_Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
