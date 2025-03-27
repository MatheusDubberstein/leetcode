package solutions

func RemoveDuplicates(nums []int) int {
	index := 0
	for indexArray, x := range nums {
		nextIndex := indexArray + 1
		if (indexArray + 1) < len(nums) {
			nextNumber := nums[nextIndex]
			if nextNumber != x {
				nums[index] = x
				index++
			}
		} else {
			nums[index] = x
			index++
		}

	}
	nums = nums[:index]
	return len(nums)
}

// func main() {
// 	nums1 := []int{1, 1, 2}
// 	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	result1 := removeDuplicates(nums1)
// 	result2 := removeDuplicates(nums2)
// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
