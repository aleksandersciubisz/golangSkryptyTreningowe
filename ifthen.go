package main

import "fmt"

func main() {

	var x float64 = 0
	var result string

	if x < 0 {
		result = "less than zero"
	} else if x == 0 {
		fmt.Println("x is exactly zero")
	} else {
		result = "Greather than zero"
	}

	fmt.Println(result)
}
