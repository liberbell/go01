package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	poodle := Dog{}
	fmt.Println(poodle)
}
