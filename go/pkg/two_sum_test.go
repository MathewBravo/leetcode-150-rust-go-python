package pkg

import "testing"

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	expect := []int{1, 2}
	pass := true
	for i, num := range expect {
		if num != result[i] {
			pass = false
		}
	}
	if !pass {
		t.Errorf("Two Sum: Expected %v, Got %v", expect, result)
	}
}
