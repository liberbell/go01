package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h hello) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from the Go web server!</h1>")
}

func main() {
	var h hello
	err := http.ListenAndServe("localhost:4000", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
