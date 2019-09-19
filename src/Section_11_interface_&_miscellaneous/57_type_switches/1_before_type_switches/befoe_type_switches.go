package main

import "fmt"

func testValue(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("string value : ", v)
		return
	}
	if v, ok := x.(int); ok {
		fmt.Println("int value : ", v)
		return
	}
	if v, ok := x.(bool); ok {
		fmt.Println("bool value : ", v)
		return
	}
	if v, ok := x.(float32); ok {
		fmt.Println("float32 value : ", v)
		return
	}
	if v, ok := x.(float64); ok {
		fmt.Println("float64 value : ", v)
		return
	}
	fmt.Println("No match any")
}

func main() {
	// testValue()
	testValue("sas")
	testValue(234)
	testValue(true)
	testValue(234.23) // default type of xx.xx is float64
	testValue(float32(34.23))
}
