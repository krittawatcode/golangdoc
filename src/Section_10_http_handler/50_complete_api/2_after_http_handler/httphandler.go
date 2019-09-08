package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// type myHandler func(http.ResponseWriter, *http.Request)

// func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	m(w, r)

// }

func inventory(w http.ResponseWriter, r *http.Request) {

	// r.URL /items
	// r.URL /prince?item=apple
	log.Println("r.url : ", r.URL.Path)

	switch r.URL.Path {
	case "/items":
		w.Write([]byte("handle items"))
	case "/price":
		w.Write([]byte("handle price"))
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry, no such page: %s", r.URL.Path)
	}

	// w.Write([]byte("Hello 2"))
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", http.HandlerFunc(inventory))

}
