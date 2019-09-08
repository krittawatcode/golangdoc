package main

import "fmt"

// struct (record)
type Person struct {
	name string
	age  int
}

/* anonymous field */
type Employee struct {
	Person
	designation string
}

func main() {
	filicity := Employee{
		Person{"Filicity", 23},
		"Programmer",
	}

	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.name)

}
