package pkg

import "testing"

func TestTrap(t *testing.T) {
	result := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	if result != 6 {
		t.Errorf("Trap Rain Water: Expected 6, Got %d", result)
	}
}
