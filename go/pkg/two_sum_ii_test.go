package pkg

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	expected := []int{1, 2}
	result := twoSumII(numbers, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	numbers = []int{2, 3, 4}
	target = 6

	expected = []int{1, 3}
	result = twoSumII(numbers, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	numbers = []int{-1, 0}
	target = -1

	expected = []int{1, 2}
	result = twoSumII(numbers, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
