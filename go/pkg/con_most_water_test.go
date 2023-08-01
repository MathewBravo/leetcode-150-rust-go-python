package pkg

import "testing"

func TestMaxArea(t *testing.T) {
	test := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	result := 49
	if result != test {
		t.Errorf("maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7] expected %d, returned %d)", result, test)
	}
}
