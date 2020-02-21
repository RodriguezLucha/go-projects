package main

import "fmt"

type animal interface {
	speak() string
}
type dog struct{}
type cat struct{}

func main() {
	doggy := dog{}
	catty := cat{}

	printGreeting(doggy)
	printGreeting(catty)
}

func printGreeting(a animal) {
	fmt.Println(a.speak())
}

func (dog) speak() string {
	return "woof woof"
}

func (cat) speak() string {
	return "meow meow"
}
