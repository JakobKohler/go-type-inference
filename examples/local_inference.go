package main

import "fmt"

func local_inference() {
	x := 42            // inferred as int
	y := 3.14          // inferred as float64
	z := complex(2, 3) // inferred as complex128

	fmt.Printf("x: %T, y: %T, z: %T\n", x, y, z)
}
