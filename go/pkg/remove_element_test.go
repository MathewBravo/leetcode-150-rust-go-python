package pkg

import "testing"

func TestRemoveElement(t *testing.T) {
	result := removeElement([]int{3, 2, 2, 3}, 3)
	if result != 2 {
		t.Errorf("Remove Element: Expected %d, Got %d", 2, result)
	}
}
