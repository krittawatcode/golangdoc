package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		Array = collection of data that arranged in Ram
			- var a [4]int <-- 1 variable
			- size of array effect type of array
				- [2]int{} != [3]int{}
			- copy by value
	*/

	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}

	fmt.Println(len(fruits), fruits)

	twoSlots := [2]int{}
	threeSlots := [3]int{}

	fmt.Println("Two slot\t", reflect.TypeOf(twoSlots))
	fmt.Println("Three slot\t", reflect.TypeOf(threeSlots))

	// Part 2
	fruits2 := [...]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits2), fruits2)

	// Part 3
	/*
		----- Gotcha -----
		{ INDEXx : value, INDEXy : value}
	*/
	fruits3 := [...]string{
		4: "apple",
		2: "banana",
		1: "papaya",
		3: "orange",
		0: "mango",
	}
	fmt.Println(len(fruits3), fruits3, fruits3[4])

	// Part 4 : compare array
	fmt.Println([2]int{1, 2} == [2]int{1, 2})
	fmt.Println([3]int{1, 2, 3} == [3]int{1, 3, 2})

	fmt.Println("----- Compare Dupplicate Array -----")
	dubFruits := fruits
	dubFruits[0] = "Watermelon"
	fmt.Println(fruits)
	fmt.Println(dubFruits)
	// fruits dont change when change value in dubFruit -> copy by value (!= copy by ref)

	// Part 5 : pass by ref
	fmt.Println("----- Compare Dubplicate Array : by ref -----")
	ptrFruits := &fruits
	ptrFruits[0] = "Strawberry"
	fmt.Println(fruits)
	fmt.Println(*ptrFruits)

	// Part 6 : Array[][]
	a := [2][3]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Println(a)
}
