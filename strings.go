package main

import (
	"fmt"
	//"strconv"
	"strings"
)

func main() {
	str1 := "Implicit"
	var str2 string = "Explicit"
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Printf("Implicit: %v\n Explicit: %v\n",
		str1, str2)
	fmt.Printf("typ str1 %T\n", str1)

	fmt.Println("pokaz string package")
	fmt.Println(strings.ToUpper(str2))

	lvalue := "Hello"
	lValue := "hello"
	fmt.Println("Equal?", (lValue == lvalue))
	fmt.Println("Equal? Non-case sensitive",
		strings.EqualFold(lValue, lvalue))
}
