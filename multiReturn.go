package main

import "fmt"

func main() {

	n1, l1 := FullName("Zarak", "Blombasa")
	fmt.Printf("Fullname: %v, number of characters: %v,",
		n1, l1)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	lenght := len(full)
	return full, lenght
}
