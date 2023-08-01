package pkg

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	test := MajorityElement([]int{3, 2, 3})
	expect := 3
	if expect != test {
		t.Errorf("Majority Element failed. Expected: %d, Got: %d", expect, test)
	}
}
