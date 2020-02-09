package main

import (
	"fmt"
	"math"
)

func main() {
	a := []float64{5, 6}
	b := []float64{5, 7}
	rse, _ := get_rse(normalise(a), normalise(b))
	fmt.Printf("%v", rse)
}

// Get pointwise root squared error
// Input values should be sorted and between 0 and 1
func get_rse(a []float64, b []float64) (rse float64, err error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("need same number of points to calculate error")
	}

	for i := range a {
		rse += math.Pow(a[i]-b[i], 2)
	}

	return math.Sqrt(rse / float64(len(a))), nil
}

// Pre: points is sorted, e.g. 3, 15, 67
func normalise(points []float64) []float64 {
	normalised := make([]float64, len(points))
	duration := points[len(points)-1]
	for i := range points {
		normalised[i] = points[i] / duration
	}

	return normalised
}
