package main

import "fmt"

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	defer fmt.Println("Statement1")
	defer fmt.Println("Statement2")
	defer fmt.Println("Statement3")
	defer fmt.Println("Statement4")
	defer fmt.Println("Statement5")
}
