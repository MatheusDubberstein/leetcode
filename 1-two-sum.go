// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int)
	var indexResult []int
	for index, x := range nums {
		indexTarget, ok := numbers[target-x]
		if ok {
			indexResult = []int{indexTarget, index}
		}
		numbers[x] = index
	}
	return indexResult
}

func main() {
	var result []int

	nums1 := []int{2, 7, 11, 15}
	result = twoSum(nums1, 9)
	fmt.Println(result)

	nums2 := []int{3, 2, 4}
	result = twoSum(nums2, 6)
	fmt.Println(result)

	nums3 := []int{3, 3}
	result = twoSum(nums3, 6)
	fmt.Println(result)
}
