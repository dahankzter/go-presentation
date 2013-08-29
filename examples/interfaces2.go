package main

import "fmt"

// START OMIT
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

// END OMIT

// Just for kicks :)
type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
