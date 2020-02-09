package server

import (
	"fmt"
	"math"
)

func main() {
	// Straw, be-rry. Cheese-cake.
	attempt := []int{1000, 1200, 2000, 2200}

	sigs := [][]int{{800, 1100, 2100, 2200},
		{1100, 1200, 1900, 2250},
		{100, 100, 100, 100}}

	fmt.Printf("Matches majority? %v", MatchesMajority(attempt, sigs))
}

func MatchesMajority(attempt []int, sigs [][]int) bool {
	normAttempt := normalise(millisToSecs(attempt))

	matched := 0
	for _, sig := range sigs {
		normSig := normalise(millisToSecs(sig))
		if matches(normAttempt, normSig) {
			matched++
		}
	}

	return float64(matched) >= float64(len(sigs))/2
}

// Pre: both are normalised
func matches(a []float64, b []float64) bool {
	threshold := 0.1
	rse, _ := getRse(a, b)
	return rse < threshold
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
	for i, point := range points {
		normalised[i] = point / duration
	}

	return normalised
}

func millisToSecs(millis []int) []float64 {
	secs := make([]float64, len(millis))
	for i, milli := range millis {
		secs[i] = float64(milli) / 1000
	}

	return secs
}
