package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Wooof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Miau!"
}

type Cow struct {
}

func (k Cow) Speak() string {
	return "muuuu!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
