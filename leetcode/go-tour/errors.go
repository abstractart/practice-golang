package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

// Sqrt ...
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z0 := Newton(1, x)
	zi := z0
	const delta = 0.0000001

	for math.Abs(x-math.Pow(zi, 2)) > delta {
		zi = Newton(zi, x)
	}

	return zi, nil
}

// Newton formula
func Newton(zi float64, x float64) float64 {
	return zi - ((zi*zi)-x)/(2*zi)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
