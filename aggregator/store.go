package main

import "github.com/Tanmoy095/Toll-Calculator.git/types"

type MemoryStore struct {
}

func (s *MemoryStore) Insert(distance types.Distance) error {
	return nil

}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}
