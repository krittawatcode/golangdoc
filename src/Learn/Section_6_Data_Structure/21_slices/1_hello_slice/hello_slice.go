package main

import "fmt"

func main() {
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	myFavor := fruits[1:4]                           // start at 1 end exclude 4
	fmt.Println(len(myFavor), cap(myFavor), myFavor) // 3 4 [banana papaya orange]
	yourFavor := myFavor
	yourFavor[0] = "guava"
	fmt.Println(len(yourFavor), cap(yourFavor), yourFavor) // 3 4 [guava papaya orange]
	fmt.Println(len(myFavor), cap(myFavor), myFavor)       // 3 4 [guava papaya orange]
}
