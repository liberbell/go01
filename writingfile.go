package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go!"

	file, err := os.Create("./fromstring.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with file of %v characters\n", ln)

	bytes := []byte(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
