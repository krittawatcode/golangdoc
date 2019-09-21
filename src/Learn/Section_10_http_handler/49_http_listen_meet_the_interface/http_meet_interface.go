package main

import (
	"log"
	"net/http"
	"os"
)

type myHandler func()

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", myHandler(func() {

	}))
}
