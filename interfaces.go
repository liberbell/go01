package main

import "fmt"

type Animal interface {
	Speak(volume int) string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)
}
