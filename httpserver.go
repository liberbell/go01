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
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
