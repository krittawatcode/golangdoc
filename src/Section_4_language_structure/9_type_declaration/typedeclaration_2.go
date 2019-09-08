package main

import "fmt"

// *** Type declaration ***
// type YourTypeName float64
//  ^         ^         ^
// type     name   underlying-type

type KG float64
type LB float64

func (lb LB) toKG() KG {
	return KG(lb / 2.2046226128)
}

func main() {

	k := KG(3)
	b := LB(3)

	fmt.Println(k > b.toKG())
}
