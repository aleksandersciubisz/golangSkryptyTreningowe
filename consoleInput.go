package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("enter number:")
	str1, _ := reader.ReadString('\n')
	f, err := strconv.ParseFloat(str1, 64)
	fmt.Println(str)
	fmt.Println(f)
	if err != nil {
		fmt.Println("error!")
		fmt.Println(err)
	}

}
