package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poodle := Dog{"poodle", 20}
	fmt.Println(poodle)
}
