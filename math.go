package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	f1, f2, f3 := 12.3, 45.4, 68.77
	int2Sum := f1 + f2 + f3
	fmt.Println("Floating sum: ", int2Sum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.5)
	b2.SetFloat64(3.6)
	b3.SetFloat64(77.11)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("big sum: %10g\n", &bigSum)

}
