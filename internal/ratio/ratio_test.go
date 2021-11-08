package ratio

import (
	"testing"
)

func TestRatio(t *testing.T) {

	// test: once-ed configuration
	table := []struct {
		count    float64
		ratio    float64
		expected int
	}{
		{100, 6.67, 15},
		{100, 50, 2},
		{100, 20, 5},
		{300, 20, 15},
	}

	// table: iterate test cases
	for i, tc := range table {

		// table: execute
		actual := Calculate(tc.count, tc.ratio)

		// table: validate expectations
		if tc.expected != actual {
			t.Errorf("[%d] expected=%v actual=%v",
				i, tc.expected, actual)
		}

	}

}
