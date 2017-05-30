package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
	case 1:
		result = "wypada jedynka"
	case 2:
		result = "wypada dwójka"
	case 3:
		result = "wypada trójka"
	default:
		result = "wypadło coś innego niż 1,2 i 3"
	}
	fmt.Println(result)

	x := dow / 3
	switch {
	case x < 0:
		fmt.Println("X mniejsze niż zero")
	case x > 0:
		fmt.Println("X większe niż zero")
	case x == 0:
		fmt.Println("X równe zero")
	}
}
