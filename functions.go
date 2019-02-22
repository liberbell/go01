package main

import "fmt"

func main() {
	dosomething()
	sum := AddValues(23, 54)
	fmt.Println("Sum: ", sum)
}

func dosomething() {
	fmt.Println("Doing Something")
}

func AddValues(value1 int, value2 int) int {
	return value1 + value2
}
