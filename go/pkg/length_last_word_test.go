package pkg

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	test := LengthOfLastWord("Hello World")
	expect := 5
	if test != expect {
		t.Errorf("Length of last word failed. Expected: %d, Got: %d", expect, test)
	}
}
