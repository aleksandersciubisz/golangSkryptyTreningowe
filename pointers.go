package main

import "fmt"

func main() {

	var p *int

	if p != nil {
		fmt.Println("pointer value is:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 44
	p = &v

	if p != nil {
		fmt.Println("pointer value is:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 43.99
	pointer1 := &value1
	fmt.Println("value1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("value1 divided:", *pointer1)
	fmt.Println("value1 undivided:", value1)
	fmt.Println("zonk! value 1 zostało podzielone")
	fmt.Println("przez odwołanie do pointera")

}
