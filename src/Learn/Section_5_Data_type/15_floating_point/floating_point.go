package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println(x, -x, 1/x, x/x)   // 0 -0 +Inf NaN
	fmt.Println(math.IsNaN(x / x)) // true
}
