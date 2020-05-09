// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)

	maximum := candies[0]

	for i := 1; i < n; i++ {
		if candies[i] > maximum {
			maximum = candies[i]
		}
	}

	result := make([]bool, n)

	for i := 0; i < n; i++ {
		if (candies[i] + extraCandies) >= maximum {
			result[i] = true
		}
	}

	return result
}

func main() {
	var candies = []int{2, 3, 5, 1, 3}
	extraCandies := 3
	result := kidsWithCandies(candies, extraCandies)
	fmt.Println(result)
}
