package main

import (
	"fmt"
	"stringutil"
)

func main() {
	n1, l1 := stringutil.FullName("Zaphad", "Beeblebrox")
	fmt.Printf("Fullname: %v, Number of chars: %v\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, Number of chars: %v\n", n2, l2)
}
