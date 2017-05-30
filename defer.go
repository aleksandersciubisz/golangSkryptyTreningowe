package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	defer fmt.Println("statement 1")
	defer fmt.Println("statement 2")

	myFunc()

	defer fmt.Println("statement 3")
	fmt.Println("undefered statement")

	x := 1000
	defer fmt.Println("Value od x:", x)
	x++
	fmt.Println("Value od x:", x)

}

func myFunc() {

	defer fmt.Println("Defered in the function")
	fmt.Println("Non defered in the function")

}
