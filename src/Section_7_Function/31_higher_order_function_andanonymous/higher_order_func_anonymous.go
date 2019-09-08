package main

import (
	"log"
	"time"
)

func createFib() func(int) []int {
	fList := []int{0, 1, 1, 2, 3, 5}
	return func(n int) []int {
		if n > len(fList) {
			for n > len(fList) {
				lastIndex := len(fList) - 1
				fList = append(fList, fList[lastIndex]+fList[lastIndex-1])
			}

		}
		return fList[:n]
	}
}

func profileTime(f func(int) []int) func(int) []int {
	return func(a int) []int {
		start := time.Now()
		result := f(a)
		log.Printf("call with %d tooks %d ns", a, time.Now().Sub(start))
		return result
	}
}

func main() {

	fib := createFib()
	fib = profileTime(fib)
	fib(5) // 0 1 1 2 3
	fib(6) // 0 1 1 2 3 5

	fib(60000000)
	fib(60000) // ****** this one call after we extend fList so we can get data in 0 ns
}
