package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last2 := [2]int{0, 1}

	return func() int {
		val := last2[0] + last2[1]

		last2[0] = last2[1]
		last2[1] = val

		return last2[1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
