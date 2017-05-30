package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Minimal writing in Go!"

	file, err := os.Create("./mnml.txt")

	io.WriteString(file, content)

	fmt.Println("done")
	file.Close()
	if err != nil {
		fmt.Println("error")
	}
}
