package main

import (
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"

	file, err := os.Create("./fromstring.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
