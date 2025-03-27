package main

import (
	"fmt"
	"leetcode-go/solutions"
)

func main() {
	fmt.Println("Problem 1: Two Sum")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := solutions.TwoSum(nums, target)
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Ouput:  %v\n\n", result)
}
