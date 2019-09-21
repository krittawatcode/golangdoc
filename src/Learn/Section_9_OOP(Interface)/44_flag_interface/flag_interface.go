package main

import (
	"flag"
	"fmt"
	"strconv"
)

var romanMap = map[string]int{
	"i":   1,
	"ii":  2,
	"iii": 3,
	"iv":  4,
	"v":   5,
}

type RomanAge struct {
	age string
}

func (r RomanAge) String() string {
	return strconv.Itoa(romanMap[r.age])
}

func (r *RomanAge) Set(value string) error {
	r.age = value
	return nil
}

func flagRomanAge() *RomanAge {
	romanAge := RomanAge{}
	flag.Var(&romanAge, "age", "help message for flagname")
	return &romanAge
}

var name = flag.String("name", "anonymous", "your name") // input => flag name, default value, help message // return => address of that flag name

var age = flagRomanAge()

func main() {

	/* Important!!!!
	run this code with command line
	go run '.\src\Section_9_OOP(Interface)\44_flag_interface\flag_interface.go' -name Filicity
	*/

	flag.Parse()
	fmt.Println(name)  // return address
	fmt.Println(*name) // return value

	/* run without flag name */
	fmt.Println(*name)

	fmt.Println("--------------------")
	/* last part */
	fmt.Println(*age)
}
