package main

import (
	"fmt"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

type Aggregator interface {
	AggregateDistance(types.OBUData) (float64, error)
}
type storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store storer
}

func NewInvoiceAggregator(store storer) *InvoiceAggregator {
	return &InvoiceAggregator{store: store}
}

func (i *InvoiceAggregator) AggregateDistance(data types.Distance) (float64, error) {
	fmt.Println("Processing and inserting distance in the storage", data)
	return i.store.Insert(data)
}
