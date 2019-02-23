package main

import "fmt"

func main() {
	n1, l1 := FullName("Zaphad", "Beeblebrox")
	fmt.Printf("Fullname: %v, Number of chars: %v\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, Number of chars: %v\n", n2, l2)
}
