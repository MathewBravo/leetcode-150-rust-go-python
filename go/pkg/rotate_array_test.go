package pkg

import "testing"

func TestRotate(t *testing.T) {
	t_slice := []int{1, 2, 3, 4, 5, 6, 7}
	expect_slice := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(t_slice, 3)
	pass := true
	for i, num := range expect_slice {
		if t_slice[i] != num {
			pass = false
		}
	}
	if !pass {
		t.Errorf("Rotate Array Failed: Expected %v, Got %v", expect_slice, t_slice)
	}
}
