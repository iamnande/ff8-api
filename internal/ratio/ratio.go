package ratio

import (
	"math"
)

// Calculate is a function that calculates the nearest whole number of the ratio
// of count to ratio.
func Calculate(count, ratio float64) int {

	// ratio: calculate raw float
	raw := count / ratio

	// ratio: round to the nearest whole number
	value := math.Round(raw*10) / 10

	// ratio: return calculated number
	return int(value)

}
