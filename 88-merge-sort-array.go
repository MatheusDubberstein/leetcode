package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1Slice := nums1[0:m]
	nums2Slice := nums2[0:n]
	result := append(nums1Slice, nums2Slice...)
	sort.Ints(result)
	return result
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	result := merge(nums1, 3, nums2, 3)
	fmt.Println(result)
}
