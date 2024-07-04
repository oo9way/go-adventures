package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func hello2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello2", hello2)
	http.ListenAndServe(":8090", nil)
}
