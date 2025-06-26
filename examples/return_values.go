package main

import "fmt"

func getNameAndAge() (string, int) {
	return "Alice", 30
}

func inferFromReturnValue() {
	name, age := getNameAndAge() // inferred as string and int
	fmt.Printf("name: %T, age: %T\n", name, age)
}
