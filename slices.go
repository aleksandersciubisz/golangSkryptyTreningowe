package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = []string{
		"red",
		"green",
		"blue",
		"yellow",
	}

	fmt.Println(colors)
	fmt.Println(colors[2])

	colors = append(colors, "purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	colors = append(colors[1:3])
	colors = append(colors, "pink")

	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println("after sorting:", colors)

}
