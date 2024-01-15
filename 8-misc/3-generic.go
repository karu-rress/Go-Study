package main

import "fmt"

// Type concept
type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) (s V) {
	for _, v := range m {
		s += v
	}
	return
}

func Generic() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
