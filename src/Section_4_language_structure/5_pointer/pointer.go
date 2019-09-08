package main

import (
	"fmt"
)

func main() {

	x := 70
	var p *int // declare p is ref type of int
	p = &x     // p =  x location's
	// fmt.Println(x)
	fmt.Println("&x", &x) // show location in memory
	fmt.Println("p", p)
	fmt.Println("&p", &p)
	fmt.Println("value of *p", *p) // de reference pointer -> value of x
	fmt.Println(*(&x))

}
