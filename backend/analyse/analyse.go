package main

import (
	"fmt"
	"math"
)

func main() {
	rse, _ := get_rse([]float64{0.5, 0.6}, []float64{0.5, 0.7})
	fmt.Printf("%v", rse)
}

// Get pointwise root squared error
// Input values should be sorted and between 0 and 1
func get_rse(a []float64, b []float64) (rse float64, err error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Need same number of points to calculate error")
	}

	for i := range a {
		rse += math.Pow(a[i]-b[i], 2)
	}

	return math.Sqrt(rse / float64(len(a))), nil
}
