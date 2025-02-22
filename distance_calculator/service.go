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
	points [][]float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		points: make([][]float64, 0),
	}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0

	if len(s.points) > 0 {
		prevPoint := s.points[len(s.points)-1]
		distance = calculateDistance(prevPoint[0], prevPoint[1], data.Latitude, data.Longitude)

	}
	s.points = append(s.points, []float64{data.Latitude, data.Longitude})
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
