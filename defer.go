package main

import "fmt"

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	defer fmt.Println("Statement1")
	defer fmt.Println("Statement2")

	myFunc()

	defer fmt.Println("Statement3")
	defer fmt.Println("Statement4")
	fmt.Println("Undefered Statement")
}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
