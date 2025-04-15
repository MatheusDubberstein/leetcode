// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.
// Example 1:
//
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:
//
// Input: height = [1,1]
// Output: 1
package solutions

func ContainerWithMostWater(height []int) int {
	leftIndex := 0
	rightIndex := 0
	maxArea := 0
	for index, currentHeight := range height {
		area := 0
		currIndex := index - 1
		leftHeight := height[leftIndex]
		rightHeight := height[rightIndex]
		if currentHeight > rightHeight {
			area = rightHeight*currIndex - rightIndex
		} else if currentHeight > leftHeight {
			area = leftHeight*currIndex - leftIndex
		}
		if leftHeight > currentHeight {
			leftIndex = currIndex
		}
		if rightHeight < currentHeight {
			rightIndex = currIndex
		}
		if area > maxArea {
			maxArea = area
		}

	}
	return maxArea
}
