package solutions

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//
// Example 1:
//
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:
//
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

//	func checkNextJump(nums []int, index int) bool {
//		number := nums[index]
//		var nextIndex int
//		if number > 1 {
//			nextIndex = index + (number - 1)
//		} else {
//			nextIndex = index + number
//		}
//
//		fmt.Println(nextIndex)
//		if nextIndex == (len(nums)-1) || index == (len(nums)-1) {
//			return true
//		}
//
//		if number == 0 {
//			return false
//		}
//
//		if nextIndex < (len(nums) - 1) {
//			return checkNextJump(nums, nextIndex)
//		}
//		return true
//	}
//
//	func canJump(nums []int) bool {
//		return checkNextJump(nums, 0)
//	}
//
// FIX: Remove this solution and think about solution
func CanJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}

// func main() {
// 	var result bool
//
// 	nums1 := []int{2, 3, 1, 1, 4}
// 	result = canJump(nums1)
// 	fmt.Println(result)
//
// 	nums2 := []int{3, 2, 1, 0, 4}
// 	result = canJump(nums2)
// 	fmt.Println(result)
//
// 	nums3 := []int{2, 0, 0}
// 	result = canJump(nums3)
// 	fmt.Println(result)
// }
