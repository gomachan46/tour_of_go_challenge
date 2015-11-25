package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 0.001
	for {
		next := z - (z*z-x)/2*z
		if math.Abs(next-z) < delta {
			break
		}
		z = next
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
