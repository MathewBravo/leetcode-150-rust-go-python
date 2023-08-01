package pkg

import "testing"

func TestRemoveDuplicates2(t *testing.T) {
	test_var := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	RemoveDuplicates2(test_var)
	expect := []int{0, 0, 1, 1, 2, 3, 3}
	for i, num := range expect {
		if num != test_var[i] {
			t.Errorf("Remove Duplicate 2 broken, Expected %v, Got %v", expect, test_var)
		}
	}

}
