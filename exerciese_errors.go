package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	//return fmt.Sprintf(e)
}

func Sqrt(x float64) (float64, error) {
	// return 0, nil
	if x < 0 {
		return float64(0), ErrNegativeSqrt(x)
	} else {
		z := x / 2
		for delta := 1.0; delta > 1e-12; delta = math.Abs(z*z - x) {
			z -= (z*z - x) / (2 * z)
			fmt.Println(z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
