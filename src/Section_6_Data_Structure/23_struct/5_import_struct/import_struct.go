package main

import (
	"fmt"
	"hr"
)

func main() {
	// when you import Go will recommend to use property tag
	filicity := hr.Employee{
		Person:      hr.Person{"Filicity", 23},
		Designation: "Programmer",
	}

	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.Name)

}
