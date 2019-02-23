package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 49, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()
}
