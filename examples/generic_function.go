package main

import (
	"fmt"
)

func Map[S ~[]E, E any, R any](input S, f func(E) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func runGenericFunctionInference() {
	type Names []string
	var names Names = Names{"Alice", "Bob", "Charlie"}

	// Without type inference
	var toLengths func(string) int = func(s string) int {
		return len(s)
	}
	var lengths []int = Map[Names, string, int](names, toLengths)
	fmt.Println("Without type inference:", lengths)

	// With type inference
	toLengths2 := func(s string) int {
		return len(s)
	}
	lengths2 := Map(names, toLengths2)
	fmt.Println("With type inference:   ", lengths2)
}
