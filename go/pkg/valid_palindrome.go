package pkg

import "unicode"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	left, right := 0, len(s)-1
	for left <= right {
		if unicode.IsLetter(rune(s[right])) || unicode.IsDigit(rune(s[right])) {
			left++
			continue
		} else {
			right--
			continue
		}
		if unicode.IsLower(rune(s[left])) != unicode.IsLower(rune(s[left])) {
			return false
		}
		left++
		right++
	}
	return true
}
