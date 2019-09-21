package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe("", nil)
}
