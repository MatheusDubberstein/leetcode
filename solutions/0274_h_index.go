// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.
//
// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
//
// Example 1:
//
// Input: citations = [3,0,6,1,5]
// Output: 3
// Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
// Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
// Example 2:
//
// Input: citations = [1,3,1]
// Output: 1
package solutions

import (
	"sort"
)

func HIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	hIndex := 1
	for _, x := range citations {
		if x >= hIndex {
			hIndex++
		}
	}
	return hIndex - 1
}

// func main() {
// 	var result int
// 	numbers1 := []int{3, 0, 6, 1, 5}
// 	result = hIndex(numbers1)
// 	fmt.Println(result)
//
// 	numbers2 := []int{1, 3, 1}
// 	result = hIndex(numbers2)
// 	fmt.Println(result)
// }
