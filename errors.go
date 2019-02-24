package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("Filename.ext")
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}
}
