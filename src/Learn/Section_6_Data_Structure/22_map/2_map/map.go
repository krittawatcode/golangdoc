package main

import "fmt"

func main() {
	items := map[string]int{
		"pen":    4,
		"pencil": 5,
		"a":      0,
		"b":      0,
		"c":      0,
	}

	// map don't be in order
	for k, v := range items {
		fmt.Println(k, v)
	}
	// map don't has address
}
