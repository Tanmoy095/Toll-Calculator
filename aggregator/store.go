package main

import (
	"fmt"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}
func (s *MemoryStore) Insert(distance types.Distance) error {
	s.data[distance.OBUID] += distance.Value
	return nil

}

func (s *MemoryStore) Get(id int) (float64, error) {
	dist, ok := s.data[id]
	if !ok {
		return 0, fmt.Errorf("could not found distance for obu id  %d", id)
	}
	return dist, nil

}
