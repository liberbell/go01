package main

import "fmt"

func main() {
	dosomething()
	sum := AddValues(23, 54)
	fmt.Println("Sum: ", sum)

	sum = AddAllValues(10, 23, 44)
	fmt.Println("New sum: ", sum)
}

func dosomething() {
	fmt.Println("Doing Something")
}

func AddValues(value1 int, value2 int) int {
	return value1 + value2
}

func AddAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
