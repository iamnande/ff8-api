package calculator

//go:generate mockgen -destination=./mocks/calculator.go -package mocks github.com/iamnande/ff8-api/internal/calculator Calculator

import (
	"context"

	"github.com/iamnande/ff8-api/internal/datastore"
	"github.com/iamnande/ff8-api/internal/ratio"
)

// Response is the defined response of all calculations.
type Response struct {
	Name     string
	Card     string
	Quantity int
}

// Calculator interface describes the methods that all Calculator providers
// must satisfy.
type Calculator interface {

	// CardMagicRatio is the method definition that calculates the number of
	// cards, and the type of card, required in order to yield n magic.
	CardMagicRatio(
		ctx context.Context,
		name string,
		recordType datastore.RecordType,
		quantity float64,
	) (res *Response, err error)
}

// calculator is the currently used implementation of the Calculator.
type calculator struct {
	db datastore.Datastore
}

// compile time validation that the current implementation satisfies the
// defined interface.
var _ Calculator = (*calculator)(nil)

// NewCalculator creates a fresh in stance of a Calculator implementation.
func NewCalculator(db datastore.Datastore) Calculator {
	return &calculator{db: db}
}

// CardMagicRatio calculates the number of cards, and the type of card,
// required in order to yield n magic.
func (c *calculator) CardMagicRatio(
	ctx context.Context,
	name string,
	recordType datastore.RecordType,
	quantity float64,
) (res *Response, err error) {

	// cmr: initialize "done"-er channel
	done := make(chan struct{})

	// cmr: execute underlying business logic
	go func() {
		defer close(done)
		res, err = c.cardMagicRatio(name, recordType, quantity)
	}()

	// cmr: close everything out, or deadline
	select {
	case <-done:
		return res, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}

// cardMagicRatio calculates the number of cards, and the type of card,
// required in order to yield n magic.
func (c *calculator) cardMagicRatio(
	name string,
	recordType datastore.RecordType,
	quantity float64,
) (*Response, error) {

	// cmr: fetch desired item from datastore
	record, err := c.db.DescribeRecord(name, recordType)
	if err != nil {
		return nil, err
	}

	// cmr: perform ratio calculation
	return &Response{
		Name:     record.Name,
		Card:     record.CardEquivalent,
		Quantity: ratio.Calculate(quantity, record.CardMagicRatio),
	}, nil

}
