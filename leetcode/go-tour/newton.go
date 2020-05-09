package main

import (
	"fmt"
	"math"
)

// Sqrt Calculate Square
func Sqrt(x float64) float64 {
	z0 := Newton(1, x)
	zi := z0
	const delta = 0.0000001

	for math.Abs(x-math.Pow(zi, 2)) > delta {
		fmt.Println(zi)
		zi = Newton(zi, x)
	}

	return zi
}

// Newton formula
func Newton(zi float64, x float64) float64 {
	return zi - ((zi*zi)-x)/(2*zi)
}

// func main() {
// 	fmt.Println(Sqrt(9))
// }
