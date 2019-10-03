package main

import (
	"fmt"
	"math/rand"
	"time"
)

func training(ch chan bool) {
	rand.Seed(int64(time.Now().Nanosecond()))
	x := rand.Intn(1000) // even you not understand -> in conclusion it will generate new random number
	fmt.Println("training for : ", x, "milisecond")
	// time.Sleep(time.Duration(x) * time.Millisecond)
	ch <- true
}

func main() {
	done := make(chan bool)
	// training()
	go training(done) // dont know when it done? maybe before "Done" or maybe after
	// that why we should use channels to connect to the main
	trainingStatus := <-done
	fmt.Println("Done Status : ", trainingStatus)
}
