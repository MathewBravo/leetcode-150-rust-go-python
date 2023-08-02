package pkg

import "testing"

func TestPalindrome(t *testing.T) {
	result := isPalindrome("race a car")
	if !result {
		t.Errorf("Expected True Got %v", result)
	}
}
