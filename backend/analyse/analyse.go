package main

import (
	"fmt"
	"math"
)

func main() {
	// Straw, be-rry. Cheese-cake.
	rawA := []int{1000, 1200, 2000, 2200}
	rawB := []int{800, 1100, 2100, 2200}

	a := normalise(millisToSecs(rawA))
	b := normalise(millisToSecs(rawB))

	rse, _ := getRse(a, b)
	fmt.Printf("%v\n", rse)

	threshold := 0.1
	matches := rse < threshold
	fmt.Printf("Matches? %v\n", matches)
}

// Get pointwise root squared error
// Input values should be sorted and between 0 and 1
func getRse(a []float64, b []float64) (rse float64, err error) {
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

func millisToSecs(millis []int) []float64 {
	secs := make([]float64, len(millis))
	for i := range secs {
		secs[i] = float64(millis[i]) / 1000
	}

	return secs
}
