package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	pre := 0.0
	for math.Abs( z - pre) > 1e-6 {
		pre, z = z, z - (z*z - x) / (2*z)
	}
 	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
