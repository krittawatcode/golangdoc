package main

import "fmt"

func testValue(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Println("string value : ", v)
	case int:
		fmt.Println("int value : ", v)
	case bool:
		fmt.Println("bool value : ", v)
	case float32:
		fmt.Println("float32 value : ", v)
	case float64:
		fmt.Println("float64 value : ", v)
	default:
		fmt.Println("No match any")
	}
}

func main() {
	// testValue()
	testValue("sas")
	testValue(234)
	testValue(true)
	testValue(234.23) // default type of xx.xx is float64
	testValue(float32(34.23))
}
