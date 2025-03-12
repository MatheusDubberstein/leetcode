// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
//
//
// Example 1:
//
// Input: nums = [3,2,3]
// Output: 3
// Example 2:
//
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	numberMaxCount, count := nums[0], 0
	for _, num := range nums {
		if count == 0 {
			numberMaxCount = num
		}
		if numberMaxCount == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return numberMaxCount
}

func main() {
	nums1 := []int{3, 2, 3}
	nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	result1 := majorityElement(nums1)
	result2 := majorityElement(nums2)
	fmt.Println(result1)
	fmt.Println(result2)
}
