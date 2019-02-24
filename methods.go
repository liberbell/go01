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

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 49, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()
	poodle.SpeakThreeTimes()

	poodle.Sound = "Arf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
}
