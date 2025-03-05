package main

import (
	"fmt"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

const basePrice = 10.0

type Aggregator interface {
	AggregateDistance(types.Distance) error
	Calculate_Invoice(obuID int) (*types.Invoice, error)
}
type storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
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
func (i *InvoiceAggregator) Calculate_Invoice(obuID int) (*types.Invoice, error) {

	dist, err := i.store.Get(obuID)
	if err != nil {
		return nil, fmt.Errorf("obu id %d not found", obuID)

	}
	inv := &types.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalAmount:   basePrice * dist, // 10 cents per km
	}
	return inv, nil

}
