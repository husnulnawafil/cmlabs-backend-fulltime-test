package palindrome

import (
	"strings"
)

func IsPalindrome(s string) bool {
	s2 := strings.Join(strings.Split(strings.Join(strings.Split(strings.ToLower(s), " "), ""), ","), "")
	for i := 0; i < len(s2)/2; i++ {
		if s2[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}
