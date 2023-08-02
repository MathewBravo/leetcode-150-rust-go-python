package pkg

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input:    []int{3, 0, -2, -1, 1, 2},
			expected: [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
	}

	for _, test := range tests {
		result := threeSum(test.input)
		if len(result) != 0 && len(test.expected) != 0 {
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
			}
		}
	}
}
