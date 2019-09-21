package main

// import "fmt"
// import "weight"

import (
	"fmt"

	"github.com/krittawatcode/weight"
)

func main() {
	k := weight.KG(1)
	fmt.Println(weight.KgToLb(k))
}
