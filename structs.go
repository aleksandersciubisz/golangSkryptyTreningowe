package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
}

func main() {

	poodle := Dog{"poodle", 82}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight %v\n",
		poodle.Breed, poodle.Weight)
}
