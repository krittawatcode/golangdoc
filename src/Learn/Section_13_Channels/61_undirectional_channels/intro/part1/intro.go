package main

func main() {
	ch := make(chan int)

	go f(ch)
	go f2(ch)

}

func f(ch chan<- int) {
	// send only
}

func f2(ch <-chan int) {
	// receive only
}
