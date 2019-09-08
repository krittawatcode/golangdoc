package main

import "fmt"

// struct (record)
type person struct {
	name string
	age  int
}

func main() {

	/* ---- I ---- */
	// x := struct {
	// 	name string
	// 	age  int
	// }{"filicity", 23}

	/* ---- II ---- */
	// x := struct {
	// 	name string
	// 	age  int
	// }{age: 23, name: "filicity"}

	x := person{age: 23, name: "filicity"}
	y := person{age: 23}

	fmt.Println(x, x.name, x.age)
	fmt.Println(y)

}
