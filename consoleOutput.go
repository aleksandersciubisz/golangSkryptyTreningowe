package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLenght, err := fmt.Println(str1, str2, str3)
	fmt.Println(aNumber)

	if err == nil {
		fmt.Println("no error")
		fmt.Println(stringLenght)
	}

	fmt.Printf("value of aNumber: %v\n", aNumber)
	fmt.Printf("value of boolean: %v\n", isTrue)
	fmt.Printf("value of a float: %.2f\n", float64(aNumber))
	fmt.Printf("data types: %T, %T, %T\n", aNumber, isTrue, str3)

}
