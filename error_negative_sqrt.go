package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error){
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	delta := 0.001
	for {
		next := z - (z*z-x)/2*z
		if math.Abs(next-z) < delta {
			break
		}
		z = next
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}