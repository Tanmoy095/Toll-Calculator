package main

import (
	"fmt"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) (float64, error)
}
type storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store storer
}

func NewInvoiceAggregator(store storer) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) (float64, error) {
	fmt.Println("Processing and inserting distance in the storage", distance)
	err := i.store.Insert(distance)
	if err != nil {
		return 0, err

	}
	return 0.0, nil
}
