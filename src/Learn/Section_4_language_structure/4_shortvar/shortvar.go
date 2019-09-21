package main

import (
	"fmt"
	"math"
)

func main() {
	//  var
	// x, y := 1, 2
	// z, y := y, x
	// fmt.Println(x, y, z)
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%#v, %#T\n", x, x)

	x := float64(1)
	b := math.Sqrt(2)
	c := x + b
	fmt.Println(c)
}
