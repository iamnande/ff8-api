package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {

	// main: setup inputs and context
	defaultQuantity := 300.0
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	output := make(chan []*output)
	input := map[string]float64{
		"Firaga": defaultQuantity,
	}

	// main: perform calculation
	go func() {
		output <- calculate(ctx, input)
	}()

	// main: close channel and display output
	for _, item := range <-output {
		fmt.Printf("magic=%s card=%s quantity=%d\n",
			item.Magic, item.Card, item.QuantityNeeded)
	}

}

type record struct {
	Name           string  `json:"name"`
	CardEquivalent string  `json:"card_equivalent"`
	CardMagicRatio float64 `json:"card_magic_ratio"`
}

type output struct {
	Magic string

	Card           string
	QuantityNeeded int
}

var (
	// TODO: load this from db.json
	datastore = []*record{
		{
			Name:           "Firaga",
			CardEquivalent: "Hexadragon",
			CardMagicRatio: 6.67,
		},
	}
)

func calculate(ctx context.Context, input map[string]float64) []*output {

	// calculate: initialize output
	done := make(chan struct{})
	out := make([]*output, 0)

	// calculate: perform quantity conversion
	go func() {
		defer close(done)

		for magic, quantity := range input {
			for _, item := range datastore {
				if item.Name == magic {
					value := quantity / item.CardMagicRatio
					out = append(out, &output{
						Magic:          item.Name,
						Card:           item.CardEquivalent,
						QuantityNeeded: int(math.Round(value*10) / 10),
					})
				}
			}
		}
	}()

	// calculate: return output
	select {
	case <-done:
		return out
	case <-ctx.Done():
		return nil
	}

}
