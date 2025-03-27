package solutions

import (
	"regexp"
	"strings"
)

func reverseString(s string) string {
	var result string
	for _, x := range s {
		result = string(x) + result
	}
	return result
}

func IsPalindrome(s string) bool {
	alphaNumericRegex := regexp.MustCompile("[^a-zA-Z0-9]+")

	s = strings.ToLower(alphaNumericRegex.ReplaceAllString(s, ""))
	reverseS := reverseString(s)
	return reverseS == s
}

// func main() {
// 	var result bool
// 	result = isPalindrome("A man, a plan, a canal: Panama")
// 	fmt.Println(result)
//
// 	result = isPalindrome("race a car")
// 	fmt.Println(result)
//
// 	result = isPalindrome(" ")
// 	fmt.Println(result)
// }
