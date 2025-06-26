package main

import "fmt"

func trivial_inference1() {
	x := 42
	fmt.Printf("Value: %d, Type: %T\n", x, x)
}

func formatterFunction[X int | float32 | float64, Y int | float32 | float64](input X, input2 Y) (X, Y) {
	return input * 2, input2 * 3
}

func formatValues[V int | float32 | float64, B int | float32 | float64](a V, b B, formatter func(x V, y B) (V, B)) string {
	x, y := formatter(a, b)
	return fmt.Sprintf("Value 1: %v, Value 2: %v", x, y)
}

func generic_inference() {
	// Without type inference
	out1 := formatValues[int, float32](1, 2, formatterFunction[int, float32])
	// With type inference
	out2 := formatValues(1, 2, formatterFunction)

	fmt.Println(out1)
	fmt.Println(out2)
}

type Printable interface {
	toString()
}

func printInput[E Printable](input E) { input.toString() }

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) toString() {}

func showExplicitTypeArgs() {
	x := Rectangle{width: 2, height: 3}

	printInput[Rectangle](x)

}

func test[Q any](a Q, b Q) {
	fmt.Println(a, b)
}

func run() {
	test(1, 3.14)   //This doesnt cause an error, the larger type is selected
	test(1, "test") //This causes an error due to default type mismatch
}
