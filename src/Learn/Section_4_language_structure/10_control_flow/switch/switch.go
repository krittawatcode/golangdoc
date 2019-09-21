package main

import "fmt"

func main() {

	// switch in GO dont have breakpoint -> use 'fallthrough' instead
	x := 2
	switch x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default: // Default use when no condition meet!
		fmt.Println("I will used when no condition meet!")
	}

	/* --------------------------------------------------------------- */
	fmt.Println("------------ Part II ------------")

	y := 200
	switch true {
	case 1 <= y && y < 100:
		fmt.Println("deci")
	case 100 <= y && y < 1000:
		fmt.Println("centi")
	case 1000 <= y:
		fmt.Println("meter")
	default:
		fmt.Println("I will used when no condition meet!")
	}

	/* --------------------------------------------------------------- */
	fmt.Println("------------ Part III ------------")

	z := 20
	switch false {
	case 1 <= z && z < 100:
		fmt.Println("deci")
		fallthrough
	case 100 <= z && z < 1000:
		fmt.Println("centi")
		fallthrough
	case 1000 <= z:
		fmt.Println("meter")
		fallthrough
	default:
		fmt.Println("I will used when no condition meet!")
	}

	/* --------------------------------------------------------------- */
	fmt.Println("------------ Part IV : Switch String I  ------------")

	p := "b"
	switch p {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c":
		fmt.Println("C")
	default:
		fmt.Println("I will used when no condition meet!")
	}

	/* --------------------------------------------------------------- */
	fmt.Println("------------ Part V : Switch String II ------------")

	q := "f"
	switch q {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c":
		fmt.Println("C")
	case "d", "e", "f", "g":
		fmt.Println("d", "e", "f", "g")
	default:
		fmt.Println("I will used when no condition meet!")
	}

	/* --------------------------------------------------------------- */
	fmt.Println("------------ Part V : Switch String III ------------")

	r := "g" // change r to "f" and "g"
	switch r {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c":
		fmt.Println("C")
	case "d", "e", "f", "g":
		fmt.Println("d", "e", "f", "g")
		if r == "f" {
			break
		}
		fmt.Println("break not use")
	default:
		fmt.Println("I will used when no condition meet!")
	}
}
