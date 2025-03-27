// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
//
//
// Example 1:
//
// Input: s = "abc", t = "ahbgdc"
// Output: true
// Example 2:
//
// Input: s = "axc", t = "ahbgdc"
// Output: false
//
//
// Constraints:
//
// 0 <= s.length <= 100
// 0 <= t.length <= 104
// s and t consist only of lowercase English letters.

package solutions

func IsSubsequence(s string, t string) bool {
	indexS := 0
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) || len(t) == 0 {
		return false
	}
	for _, letter := range t {
		if indexS < len(s) {
			if string(letter) == string(s[indexS]) {
				indexS++
			}
		}
	}
	return indexS == len(s)
}

// func main() {
// 	var result bool
// 	result = isSubsequence("abc", "ahbgdc")
// 	fmt.Println(result)
// 	result = isSubsequence("axc", "ahbgdc")
// 	fmt.Println(result)
// 	result = isSubsequence("", "ahbgdc")
// 	fmt.Println(result)
// 	result = isSubsequence("b", "abc")
// 	fmt.Println(result)
// }
