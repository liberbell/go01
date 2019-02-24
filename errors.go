package main

import (
	"errors"
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

	myError := errors.New("My error string")
	fmt.Println(myError)
}
