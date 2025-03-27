package solutions

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:
//
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	if k == 0 || n <= 1 {
		return
	}

	temp := append([]int{}, nums[n-k:]...)
	copy(nums[k:], nums[:n-k])
	copy(nums[:k], temp)
	// fmt.Println(nums)
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7}
// 	nums2 := []int{-1, -100, 3, 99}
// 	rotate(nums, 3)
// 	rotate(nums2, 2)
// }
