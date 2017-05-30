package main

import "fmt"

func main() {
	doSomething()

	sum := addValues(1, 2)
	fmt.Println("sum:", sum)

	sum = addAllvalues(1, 3, 4, 5)
	fmt.Println("new sum:", sum)

}

func doSomething() {
	fmt.Println("do something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllvalues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
