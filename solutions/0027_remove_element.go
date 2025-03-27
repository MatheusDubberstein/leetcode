package solutions

func RemoveElement(nums []int, val int) int {
	index := 0
	for _, x := range nums {
		if x != val {
			nums[index] = x
			index++
		}
	}
	nums = nums[:index]
	return len(nums)
}

// func main() {
// 	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
// 	result1 := removeElement(nums1, 2)
//
// 	fmt.Println(result1)
//
// 	nums2 := []int{3, 2, 2, 3}
// 	result2 := removeElement(nums2, 3)
//
// 	fmt.Println(result2)
// }
