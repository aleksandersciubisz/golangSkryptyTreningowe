package main

import (
	"fmt"
)

func main() {

	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	fmt.Println(colors)
	fmt.Println(colors[2])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers[3], numbers[1])

	fmt.Println("number of colors:",
		len(colors))
}
