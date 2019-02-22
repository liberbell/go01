package main

import "fmt"

func main() {
	dosomething()
	sum := AddValues(23, 54)
	fmt.Println("Sum: ", sum)

  sum AddAllvalues(10, 23, 44)
}

func dosomething() {
	fmt.Println("Doing Something")
}

func AddValues(value1 int, value2 int) int {
	return value1 + value2
}

func AddAllValues(values ...int) int {
  sum := 0
  fmt.Printf("%T\n", values)
  return sum
}
