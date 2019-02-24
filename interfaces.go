package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
}

func (d Dog) speak() {
	return "Woof!"
}

func main() {
	poodle := Dog{}
	fmt.Println(poodle)
}
