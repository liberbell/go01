package main

import "os"

func main() {
	content := "Hello from Go!"

	file, err := os.Create("./fromstring.txt")
}

func checkError(err error) {
}
