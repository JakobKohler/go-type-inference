package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func inferComposite() {
	p := Person{"Bob", 40}                         // inferred as Person
	nums := []int{1, 2, 3}                         // inferred as []int
	settings := map[string]bool{"dark_mode": true} // inferred as map[string]bool

	fmt.Printf("p: %T, nums: %T, settings: %T\n", p, nums, settings)
}
