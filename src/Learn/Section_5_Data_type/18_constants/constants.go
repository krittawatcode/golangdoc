package main

/*
	Constants
	- complier knows it values
	- evaluation  at complie time
	- Declare constant as a group
	- lota
	- Untypeed constant
*/

import (
	"fmt"
	"reflect"
)

func main() {

	// part 1
	var x = 2
	const y = 5 - 6
	fmt.Println(x / y)

	// part 2
	const persons = 4
	toffee := 5 / persons
	cost := 2.0 / persons

	fmt.Println(toffee, reflect.TypeOf(toffee))
	fmt.Println(cost, reflect.TypeOf(cost))

	// part 3
	type Peice int
	toffee2 := Peice(5) / persons
	fmt.Println(toffee2, reflect.TypeOf(toffee2))

	// part 4 const group
	// const (
	// 	Sunday    = 1
	// 	Monday    = 2
	// 	Tuesday   = 3
	// 	Wednesday = 4
	// 	Thursday  = 5
	// 	Friday    = 6
	// 	Saturday  = 7
	// )
	// fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// part 5 const group with out iota
	// const (
	// 	Sunday = 1
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Thursday
	// 	Friday
	// 	Saturday
	// )
	// fmt.Println("const group with out iota : ", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		Sunday = iota
		// Sunday = iota + 1
		// Sunday = float64 = iota + 1.1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("const group with iota : ", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
