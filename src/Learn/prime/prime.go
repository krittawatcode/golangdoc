package prime

import (
	"fmt"
	"math"
)

var notPrimes [1000000]bool

// ------- Before tunning ------------------------
// func init() {
// 	fmt.Println("initialization in Prime Package")

// 	for i := 2; i < len(notPrimes); i++ {
// 		if notPrimes[i] {
// 			continue
// 		}
// 		notPrimes[i] = false
// 		for j := i * 2; j < len(notPrimes); j += i {
// 			notPrimes[j] = true
// 		}
// 	}
// 	fmt.Println("initialized")

// }

// ------- After tunning ------------------------

func init() {
	fmt.Println("initialization in Prime Package")
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrimes)))))
	for i := 2; i < sqrtLen; i++ {
		if notPrimes[i] {
			continue
		}
		notPrimes[i] = false
		for j := i * 2; j < len(notPrimes); j += i {
			notPrimes[j] = true
		}
	}
	fmt.Println("initialized")
}

func IsPrime(x int) bool {
	return !notPrimes[x]
}
