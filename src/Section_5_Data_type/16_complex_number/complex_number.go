package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	/* ----- Demo I ----- */
	// x := complex(7, 3)
	// fmt.Println(x)

	// y := complex128(complex(7, 3))
	// fmt.Println(y)

	// var z complex128 = complex(7, 3)
	// fmt.Println(x, z)

	/* ----- Demo II ----- */
	x := complex128(complex(3, 2))
	var y complex128 = complex(1, 7)
	fmt.Println(x + y)
	fmt.Println(x * y)
	fmt.Println(cmplx.Sqrt(-9)) // square root
	fmt.Println(cmplx.Abs(complex(3, -4)))

}
