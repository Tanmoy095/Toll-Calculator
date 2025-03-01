package main

import (
	"fmt"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}
type storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store storer
}

func NewInvoiceAggregator(store storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("Processing and inserting distance in the storage", distance)
	err := i.store.Insert(distance)
	if err != nil {
		panic(err)

	}
	return nil
}
