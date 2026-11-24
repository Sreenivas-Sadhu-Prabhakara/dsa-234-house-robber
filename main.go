package main

import "fmt"

// Rob returns the most money stealable from a row of houses when adjacent
// houses can't both be robbed. Two running values: the best including the
// previous house, and the best excluding it.
func Rob(nums []int) int {
	robbed, skipped := 0, 0 // best if previous house was robbed / wasn't
	for _, v := range nums {
		robbed, skipped = skipped+v, max(robbed, skipped)
	}
	return max(robbed, skipped)
}

func main() {
	fmt.Println(Rob([]int{1, 2, 3, 1}))
	fmt.Println(Rob([]int{2, 7, 9, 3, 1}))
}
